package search_test

import (
	"fmt"
	"testing"

	"github.com/lxkedinh/dsearch/search"
)

var fuzzySearchTests = []struct {
	name      string
	startDir  string
	query     string
	threshold float64
}{
	{"root dir level", "./testdata/search_same_dir", "ouch", 0.7},
	{"deeply nested", "./testdata/search_deeply_nested", "note", 0.4},
}

func TestFuzzyDirSearch(t *testing.T) {
	for _, test := range fuzzySearchTests {
		testName := fmt.Sprintf("fuzzy search %s test", test.name)
		t.Run(testName, func(t *testing.T) {
			results, err := search.FuzzySearch(test.startDir, test.query, 0.5)
			if err != nil {
				t.Fatalf("%s failed.\nError: %v", testName, err)
			}

			fmt.Printf("Searching for \"%s\" in starting directory \"%s\" got:\n", test.query, test.startDir)
			for _, result := range results {
				fmt.Println("-", result)
			}
		})
	}
}
