package zipfile

import (
	"path/filepath"
	"strings"
)

func getFolderNameWithoutExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func getFolderContentOutputPath(inputPath string, outputPath string) string {
	fn := getFolderNameWithoutExt(inputPath)
	return filepath.Join(outputPath, fn)
}

func getFileContentOutputPath(outputPath string, fileName string) string {
	return filepath.Join(outputPath, fileName)
}
