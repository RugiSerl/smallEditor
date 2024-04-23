package IO

import (
	"io"
	"os"
)

// extends from os.File
type file struct {
	*os.File
}

// Create a new instance of file
func ParseFile(path string) (*file, error) {
	f, err := os.Open(path)
	if err != nil {

		return nil, err
	}
	return &file{f}, nil
}

// Get the content of the file as a string
func (f *file) GetText() string {
	bytes, _ := io.ReadAll(f)
	return string(bytes)
}
