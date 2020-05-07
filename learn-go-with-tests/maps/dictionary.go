package main

// Search retrieves the value of the specified key word
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
