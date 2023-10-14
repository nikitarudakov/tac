package zipfile

import (
	"path/filepath"
	"strings"
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
