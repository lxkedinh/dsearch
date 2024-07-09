package search

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/avito-tech/normalize"
)

func Fuzzy(startDir string, query string, fuzzyThreshold float64) ([]string, error) {
	results := []string{}
	strBuilder := strings.Builder{}

	err := filepath.WalkDir(startDir, func(path string, d fs.DirEntry, err error) error {
		fileName := d.Name()
		trimmedfileName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
		if normalize.AreStringsSimilar(trimmedfileName, query, float64(fuzzyThreshold)) {
			strBuilder.WriteString(path)
			strBuilder.WriteRune(os.PathSeparator)
			strBuilder.WriteString(color.With(color.Yellow, fileName))
			results = append(results, strBuilder.String())
			strBuilder.Reset()
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
