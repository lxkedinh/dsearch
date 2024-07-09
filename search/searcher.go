package search

import (
	"io/fs"
	"path/filepath"

	"github.com/avito-tech/normalize"
)

func FuzzySearch(startDir string, query string, fuzzyThreshold float64) ([]string, error) {
	results := []string{}

	err := filepath.WalkDir(startDir, func(path string, d fs.DirEntry, err error) error {
		fileName := d.Name()
		trimmedfileName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
		if normalize.AreStringsSimilar(trimmedfileName, query, float64(fuzzyThreshold)) {
			results = append(results, fileName)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
