package filenames

import (
	"github.com/nikitarudakov/tac/groupio"
	"strings"
	"testing"
	"time"
)

func stringRepl(s string) string {
	return strings.ToUpper(s)
}

func dateRepl(ts string) string {
	t, _ := time.Parse("2006-01-02", ts)

	t = t.AddDate(0, 0, -1)

	return t.Format("2006-01-02")
}

func TestRenameItemsAtPath(t *testing.T) {

	mapper := map[string]groupio.ExprGroup{
		"Name": {GroupName: "Name", Repl: stringRepl},
		"Date": {GroupName: "Date", Repl: dateRepl},
	}

	testCases := []struct {
		path            string
		exprGroupMapper map[string]groupio.ExprGroup
	}{
		{path: "../input/renaming/test", exprGroupMapper: mapper},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.path, func(t *testing.T) {
			t.Parallel()

			err := walkInDirAndRenameFiles(test.path,
				"(?P<Name>File)_(?P<Date>.+?).xlsx",
				test.exprGroupMapper)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
