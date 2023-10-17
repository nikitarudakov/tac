package fileutils

import (
	"path/filepath"
	"strings"
)

func GetItemPathWithoutName(path string) string {
	return strings.TrimSuffix(path, filepath.Base(path))
}

func GetItemNameWithoutExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func GetItemContentOutputPath(inputPath string, outputPath string) string {
	fn := GetItemNameWithoutExt(inputPath)
	return filepath.Join(outputPath, fn)
}

func GetFileContentOutputPath(outputPath string, fileName string) string {
	return filepath.Join(outputPath, fileName)
}
