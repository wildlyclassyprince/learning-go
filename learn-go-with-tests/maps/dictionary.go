package main

import "errors"

// Dictionary mapper
type Dictionary map[string]string

// Search retrieves the value of the specified key word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
