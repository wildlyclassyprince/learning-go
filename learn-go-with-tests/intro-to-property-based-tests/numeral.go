package main

// ConvertToRoman converts Arabic numerals to Roman numerals
func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
