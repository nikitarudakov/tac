package filenames

import (
	"testing"
)

func replTest(string) string {
	return ""
}

type GroupRename struct {
	groupName string
	renameTo  string
	repl      func(string) string
}

func TestRenameItemsAtPath(t *testing.T) {
	testCases := []struct {
		path string
	}{
		{"../input/renaming/test"},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.path, func(t *testing.T) {
			t.Parallel()

			err := RenameFileWithPattern(test.path, "(?P<Name>Scoring)_(?P<Date>.+?).xlsx")
			if err != nil {
				t.Error(err)
			}
		})
	}
}
