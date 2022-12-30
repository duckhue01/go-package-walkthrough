package fs_test

import (
	"io/fs"
	"testing"
)

func TestFS(t *testing.T) {
	// fs.ReadFile()
	res, err := fs.ReadFile(nil, "file.txt")
	t.Log(res, err)
}
