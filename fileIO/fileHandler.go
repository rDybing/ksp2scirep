package fileIO

import (
	"fmt"
	"os"
)

// **************************************************** Readers ********************************************************

// LoadDMM loads a given file and returns a byte slice
func LoadSave(dir, file string) ([]byte, error) {
	f, err := os.ReadFile(dir + "/" + file)
	if err != nil {
		return f, fmt.Errorf("could not load file: %s\n%v", file, err)
	}
	return f, nil
}

// ReadDirFiles returns all files (not dirs) in a given directory and returns a string slice
func ReadDir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("could not read directory %s\n%v", dir, err)
	}
	files, err := f.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("could not read directory files: %v", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("directory %s is empty", dir)
	}
	var out []string
	for _, v := range files {
		if !v.IsDir() {
			out = append(out, v.Name())
		}
	}
	return out, nil
}

// **************************************************** Writers ********************************************************
