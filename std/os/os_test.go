package os_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func recursiveOpen(fsys fs.FS) (int, error) {

	count := 0
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil

}

func TestRecursiveOpen(t *testing.T) {
	got, err := recursiveOpen(&fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	})

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, 4) {
		t.Errorf(cmp.Diff(got, 4))
	}
}

func BenchmarkRecursiveOpen(b *testing.B) {

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

func IORedirection() {
	fmt.Fprint(os.Stdout, "asdasasd")
}

type Write struct {
}

func (w *Write) Write(p []byte) (n int, err error) {
	return 1, nil
}

func TestIORedirection(t *testing.T) {
	IORedirection()
}
