package filenames

import (
	"fmt"
	"github.com/nikitarudakov/tac/fileutils"
	"github.com/nikitarudakov/tac/groupio"
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

func handleSubexpRenaming(
	lastIndex int,
	subexpIndices []int,
	subexpNames []string,
	subexpIndex int,
	exprGroupMap groupio.ExprGroupMapper,
	filename string,
) (string, int) {
	var subexpRenamed string

	leftSubexpIndex, rightSubexpIndex := subexpIndices[subexpIndex], subexpIndices[subexpIndex+1]

	// subexp actual string for fileName
	subexpStr := filename[leftSubexpIndex:rightSubexpIndex]

	exprGroupIndex := subexpIndex / 2
	if exprGroupIndex <= len(subexpNames) {
		subexpName := subexpNames[exprGroupIndex]

		if exprGroup, ok := exprGroupMap[subexpName]; ok {
			replResp := exprGroup.Repl(subexpStr)

			subexpRenamed = filename[lastIndex:subexpIndices[subexpIndex]] + replResp

			lastIndex = subexpIndices[subexpIndex+1]

			return subexpRenamed, lastIndex
		}

		subexpRenamed = filename[lastIndex:subexpIndices[subexpIndex]] + subexpStr
		lastIndex = subexpIndices[subexpIndex+1]
	}

	return subexpRenamed, lastIndex
}

func renameFileNamesWithPattern(filename string, re *regexp.Regexp, exprGroupMap groupio.ExprGroupMapper) error {
	result := ""
	lastIndex := 0

	subexpIndices := re.FindStringSubmatchIndex(filename)
	subexpNames := re.SubexpNames()

	if len(subexpIndices) < 2 {
		return nil
	}

	for i := 2; i < len(subexpIndices); i += 2 {
		subexpRenamed, newLastIndex := handleSubexpRenaming(
			lastIndex, subexpIndices, subexpNames,
			i, exprGroupMap, filename,
		)

		lastIndex = newLastIndex
		result += subexpRenamed
	}

	result += filename[lastIndex:]

	fmt.Println(result)

	return nil
}

func walkInDirAndRenameFiles(path string, exprPattern string, exprGroupMap groupio.ExprGroupMapper) error {
	re, _ := regexp.Compile(exprPattern)

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			err := renameFileNamesWithPattern(d.Name(), re, exprGroupMap)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
