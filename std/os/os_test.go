package os_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRecursiveOpen(t *testing.T) {
	t.Parallel()

	got := func(folderName string) int {
		fsys := os.DirFS(folderName)
		count := 0
		fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if filepath.Ext(path) == ".txt" {
				count++
			}
			return nil
		})

		return count
	}("./findgo")
	t.Log(got)
	if !cmp.Equal(got, 4) {
		t.Errorf(cmp.Diff(got, 4))
	}
}

func TestOpenFile(t *testing.T) {
	got := func() string {

		_, err := os.OpenFile("text.txt", os.O_RDONLY, os.ModeAppend)
		if err != nil {

		}

		return ""
	}()

	t.Log(got)

}

func TestReadFile(t *testing.T) {
	got := func() string {
		text, err := os.OpenFile("text.txt", os.O_RDONLY, os.ModeAppend)
		if err != nil {

		}
		buf := make([]byte, 1024)
		text.Read(buf)
		return string(buf)
	}()

	t.Log(got)
}
