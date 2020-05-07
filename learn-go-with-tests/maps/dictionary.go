package main

import "errors"

// ErrNotFound when word does not exist in dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

// Dictionary mapper
type Dictionary map[string]string

// Search retrieves the value of the specified key word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add inputs new words into the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
