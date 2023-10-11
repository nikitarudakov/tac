package file

import (
	"archive/zip"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DecompressZipFile(folderPath string, outputDir string) error {
	zipFolderName := strings.TrimSuffix(filepath.Base(folderPath), filepath.Ext(folderPath))

	zipFilesListing, err := zip.OpenReader(folderPath)
	if err != nil {
		log.Fatal().Err(err).Msg("ERROR reading zip file")
		return err
	}
	defer zipFilesListing.Close()

	// create output dir
	outputPath := filepath.Join(outputDir, zipFolderName)
	if err = os.Mkdir(outputPath, os.ModePerm); err != nil {
		log.Warn().Err(err).Msg("warning! couldn't create output folder")
	}

	for _, file := range zipFilesListing.File {
		fileReader, err := file.Open()
		if err != nil {
			log.Warn().Err(err).Str("file_name", file.Name).Msg("warning! couldn't read file")
			continue
		}
		defer fileReader.Close()

		outputPath := filepath.Join(outputDir, file.Name)
		fileDest, err := os.Create(outputPath)
		if err != nil {
			log.Warn().Err(err).Msg("warning! couldn't create output folder")
		}
		defer fileDest.Close()

		log.Trace().Str("output_path", outputPath).Send()

		_, err = io.Copy(fileDest, fileReader)
		if err != nil {
			log.Error().Err(err).Msg("Error copying file contents")
			continue
		}

		fmt.Println(file.Name)
	}

	return nil
}
