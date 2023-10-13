// Package zipfile provides functionality
// to serve data from compressed files/folders
package zipfile

import (
	"archive/zip"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

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

// Unzip takes zip folder path and output path as inputs
// Reads zip folder and saves its contents to new folder in outputPath
func Unzip(inputPath string, outputPath string, isFolder bool) error {
	zipContents, err := zip.OpenReader(inputPath)
	if err != nil {
		log.Fatal().
			Str("service", "Unzip").
			Err(err).
			Msg("fatal error! failed to read zip contents")
	}
	defer zipContents.Close()

	log.Trace().
		Str("input_path", inputPath).
		Str("output_path", outputPath).
		Send()

	if err = MkOutputDir(outputPath); err != nil {
		log.Warn().
			Str("service", "Unzip").
			Err(err).
			Msg("warning! problem with creating output folder")
	}

	folderContentOutputPath := getFolderContentOutputPath(inputPath, outputPath)

	if err = MkOutputDir(folderContentOutputPath); err != nil {
		log.Warn().
			Str("service", "Unzip").
			Err(err).
			Msg("warning! problem with creating folder for content")
	}

	for _, file := range zipContents.File {
		fileReader, err := ReadCompressedFile(file)
		if err != nil {
			log.Warn().Err(err).Send()
		}

		fileContentOutputPath := getFileContentOutputPath(outputPath, file.Name)

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

func ReadCompressedFile(file *zip.File) (io.ReadCloser, error) {
	fileReader, err := file.Open()

	if err != nil {
		log.Warn().Err(err).Str("file_name", file.Name).Msg("warning! couldn't read file")
		return nil, err
	}

	defer fileReader.Close()

	return fileReader, nil
}

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		log.Warn().Err(err).Msg("warning! couldn't create file")
		return nil, err
	}

	return file, nil
}

func CopyContentToFile(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Error().Err(err).Msg("Error copying file contents")

		return err
	}

	return nil
}
