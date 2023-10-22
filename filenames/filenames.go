package filenames

import (
	"fmt"
	"github.com/nikitarudakov/tac/fileutils"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
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

func doSomethingWithGroup(string) string {
	return ""
}

func RenameFileWithPattern(path string, exprPattern string) error {
	re, _ := regexp.Compile(exprPattern)

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			result := ""
			lastIndex := 0

			submatchIndices := re.FindStringSubmatchIndex(d.Name())

			if len(submatchIndices) < 2 {
				// TODO what we do if submatchIndices length < 2
				return nil
			}

			for i := 2; i < len(submatchIndices); i += 2 {
				submatchStr := d.Name()[submatchIndices[i]:submatchIndices[i+1]]

				result += d.Name()[lastIndex:submatchIndices[i]] + doSomethingWithGroup(submatchStr)

				lastIndex = submatchIndices[i+1]
			}

			result += d.Name()[lastIndex:]
			fmt.Println(result)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
