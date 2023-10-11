package file

import (
	"archive/zip"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

func DecompressZipFile(folderPath string, outputDir string) {
	zipFolderName := strings.TrimSuffix(filepath.Base(folderPath), filepath.Ext(folderPath))

	zipFilesListing, err := zip.OpenReader(folderPath)
	if err != nil {
		log.Error().Err(err).Msg("ERROR reading zip file")
	}
	defer zipFilesListing.Close()

	// create output dir
	outputPath := filepath.Join(outputDir, zipFolderName)
	log.Trace().Str("output_path", outputPath).Send()

	if err = os.Mkdir(outputPath, os.ModePerm); err != nil {
		log.Warn().Err(err).Msg("warning! couldn't create output folder")
	}

	for _, file := range zipFilesListing.File {
		fmt.Println(*file)
	}
}
