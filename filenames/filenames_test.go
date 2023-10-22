package filenames

import "testing"

func TestRenameItemsAtPath(t *testing.T) {
	testCases := []struct {
		path string
	}{
		{"../input/renaming/test"},
	}

	groupRename := []GroupRename{
		{
			groupName: "Name",
			renameTo:  "MyFile",
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.path, func(t *testing.T) {
			t.Parallel()

			err := RenameItemsAtPath(
				test.path,
				"(?P<Name>File)_(?P<Date>.+?).xlsx",
				groupRename,
			)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
