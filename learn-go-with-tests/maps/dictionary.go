package main

// Dictionary mapper
type Dictionary map[string]string

// Search retrieves the value of the specified key word
func (d Dictionary) Search(word string) string {
	return d[word]
}
