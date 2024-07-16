package matcher

import "errors"

// TODO implementation
func Fuzzy(query string, target string, threshold float64) (bool, error) {
	return false, errors.New("not implemented")
}

// TODO implementation
func Levenshtein(word1 string, word2 string) (int, error) {
	return 0, errors.New("not implemented")
}

// TODO implementation
func Normalize(word string) (string, error) {
	return "", errors.New("not implemented")
}
