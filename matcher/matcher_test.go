package matcher_test

import (
	"fmt"
	"testing"

	"github.com/lxkedinh/dsearch/matcher"
)

var fuzzyMatchTests = []struct {
	name      string
	query     string
	target    string
	threshold float64
	expected  bool
}{
	{"strict match", "tolbr.jsx", "toolbar.jsx", 0.8, true},
	{"loose match", "Blue.mp4", "Backburner.mp4", 0.4, true},
	{"mismatch", "asdfasdfsadf", "foobar", 0.4, false},
}

func TestFuzzyMatching(t *testing.T) {
	for _, test := range fuzzyMatchTests {
		t.Run(test.name, func(t *testing.T) {
			matched, err := matcher.Fuzzy(test.query, test.target, test.threshold)

			if matched != test.expected {
				t.Fatalf("test \"%s\" got %v but expected %v match:\nQuery: %s\nTarget: %s", test.name, matched, test.expected, test.query, test.target)
			}

			if err != nil {
				t.Fatalf("error occurred during test \"%s\":\nError: %v", test.name, err)
			}
		})
	}
}

var levenshteinTests = []struct {
	word1    string
	word2    string
	distance int
}{
	{"horse", "ros", 3},
	{"intention", "execution", 5},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, test := range levenshteinTests {
		testName := fmt.Sprintf("Levenshtein distance between \"%s\" and \"%s\"", test.word1, test.word2)

		t.Run(testName, func(t *testing.T) {
			distance, err := matcher.Levenshtein(test.word1, test.word2)

			if distance != test.distance {
				t.Fatalf("Test \"%s\" got %d but expected %d", testName, distance, test.distance)
			}

			if err != nil {
				t.Fatalf("Error occured during test \"%s\":\nError: %v", testName, err)
			}
		})
	}
}

var normalizerTests = []struct {
	input    string
	expected string
}{
	{"foo#(!&#)bar.txt", "foobartxt"},
	{"184testsong%$.m%$*p4", "184testsongmp4"},
}

func TestNormalizer(t *testing.T) {
	for _, test := range normalizerTests {
		testName := fmt.Sprintf("test normalizer with \"%s\"", test.input)
		t.Run(testName, func(t *testing.T) {
			normalized, err := matcher.Normalize(test.input)

			if normalized != test.expected {
				t.Fatalf("%s got \"%s\" but expected \"%s\"", testName, &normalized, test.expected)
			}

			if err != nil {
				t.Fatalf("Error occured during %s\nError: %v", testName, err)
			}
		})
	}
}
