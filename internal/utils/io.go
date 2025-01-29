package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func ListDirs(rootPath string) ([]string, error) {
	return ListFiles(rootPath, func(entry fs.DirEntry) bool {
		return entry.IsDir()
	})
}

func ListFiles(rootPath string, filterFunc func(entry fs.DirEntry) bool) ([]string, error) {
	excluded := []string{".", ".."}

	results := make([]string, 0)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if slices.Contains(excluded, d.Name()) {
			return nil
		}

		if filterFunc == nil || filterFunc(d) {
			results = append(results, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
