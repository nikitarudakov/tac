package filenames

import (
	"github.com/nikitarudakov/tac/fileutils"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

// RenameItemWithPath renames folder/files at path p and renames it
// with newName. New name can be specified with no extensions
func RenameItemWithPath(path string, newName string) error {
	pathToFileToBeRenamed := fileutils.GetItemPathWithoutName(path)

	pathToRenamedFile := filepath.Join(pathToFileToBeRenamed, newName)

	if filepath.Ext(newName) == "" {
		pathToRenamedFile = pathToRenamedFile + filepath.Ext(path)
	}

	if err := os.Rename(path, pathToRenamedFile); err != nil {
		log.Error().Str("service", "os.Rename").Msg("ERROR! failed to rename item")
		return err
	}

	return nil
}
