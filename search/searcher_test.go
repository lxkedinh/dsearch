package search_test

import (
	"fmt"
	"testing"

	"github.com/lxkedinh/dsearch/search"
)

var fuzzySearchTests = []struct {
	name     string
	startDir string
	query    string
}{
	{"root dir level", "./testdata/search_same_dir", "foo"},
	{"deeply nested", "./testdata/search_deeply_nested", "foo"},
}

func TestFuzzyDirSearch(t *testing.T) {
	for _, test := range fuzzySearchTests {
		testName := fmt.Sprintf("fuzzy search %s test", test.name)
		t.Run(testName, func(t *testing.T) {
			results, err := search.Search(test.startDir, test.query)
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
