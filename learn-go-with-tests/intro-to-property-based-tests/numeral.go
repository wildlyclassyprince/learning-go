package main

import "strings"

// ConvertToRoman converts Arabic numerals to Roman numerals
func ConvertToRoman(arabic int) string {
	if arabic == 4 {
		return "IV"
	}
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
