// Package zipfile provides functionality
// to serve data from compressed files/folders
package zipfile

import (
	"archive/zip"
	"fmt"
	"github.com/nikitarudakov/tac/fileutils"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

// Input stores input data such as path and compression formats
type Input struct {
	path string
	zip  bool
}

// Output stores output data such as path and create boolean
// which is used while creating folder at path
type Output struct {
	path   string
	create bool
}

// SetInput sets Input fields
func (in *Input) SetInput(path string, zip bool) {
	in.path = path
	in.zip = zip
}

// SetOutput sets Output fields
func (in *Output) SetOutput(path string, create bool) {
	in.path = path
	in.create = create
}

// ValidateOutputPath validates whether folder already
// exists at o - output path and return an error if
// folder exists but user requested to create folder
func ValidateOutputPath(o string, c bool) error {
	if _, err := os.Stat(o); err != nil {
		isNotExist := os.IsNotExist(err)

		if isNotExist && c {
			return nil
		} else if isNotExist && !c {
			return err
		}

		return err
	}

	return nil
}

// MkOutputDir makes directory at output path
func MkOutputDir(outputPath string) error {
	if err := ValidateOutputPath(outputPath, true); err != nil { // add "create" flag
		log.Fatal().
			Str("service", "MkOutputDir").
			Err(err).
			Msg("fatal error! failed to validate output path")

		return err
	}

	if err := os.Mkdir(outputPath, os.ModePerm); err != nil {
		log.Warn().
			Str("service", "MkOutputDir").
			Err(err).
			Msg("warning! failed to create output folder")

		return err
	}

	return nil
}

// Decompress takes compressed folder path and output path as inputs
// Reads compressed folder and serves its contents to new folder in outputPath
func Decompress(in Input, out Output) error {
	zipContents, err := zip.OpenReader(in.path)
	if err != nil {
		log.Fatal().
			Str("service", "Decompress").
			Err(err).
			Msg("fatal error! failed to read zip contents")
	}
	defer zipContents.Close()

	log.Trace().
		Str("input_path", in.path).
		Str("output_path", out.path).
		Send()

	switch out.create {
	case true:
		if err = MkOutputDir(out.path); err != nil {
			log.Warn().
				Str("service", "Decompress").
				Err(err).
				Msg("warning! problem with creating output folder")
		}
	}

	folderContentOutputPath := fileutils.GetItemContentOutputPath(in.path, out.path)

	if err = MkOutputDir(folderContentOutputPath); err != nil {
		log.Warn().
			Str("service", "Decompress").
			Err(err).
			Msg("warning! problem with creating folder for content")
	}

	for _, file := range zipContents.File {
		fileReader, err := ReadCompressedFile(file)
		if err != nil {
			log.Warn().Err(err).Send()
		}

		fileContentOutputPath := fileutils.GetFileContentOutputPath(out.path, file.Name)

		fileDest, err := CreateFile(fileContentOutputPath)
		if err != nil {
			log.Warn().Err(err).Send()
		}
		defer fileDest.Close()

		if err = CopyContentToFile(fileDest, fileReader); err != nil {
			log.Warn().Err(err).Send()
		}

		fmt.Println(file.Name)
	}

	return nil
}

// ReadCompressedFile reads compressed file of zip format
// and return file reader along with a potential error
func ReadCompressedFile(file *zip.File) (io.ReadCloser, error) {
	fileReader, err := file.Open()

	if err != nil {
		log.Warn().Err(err).Str("file_name", file.Name).Msg("warning! couldn't read file")
		return nil, err
	}

	defer fileReader.Close()

	return fileReader, nil
}

// CreateFile creates file with a name of file stored in compressed zip folder
func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		log.Warn().Err(err).Msg("warning! couldn't create file")
		return nil, err
	}

	return file, nil
}

// CopyContentToFile copies from src to dst
// Src is the file from compressed folder whereas dst is the newly created
// file at output path
func CopyContentToFile(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Error().Err(err).Msg("Error copying file contents")

		return err
	}

	return nil
}
