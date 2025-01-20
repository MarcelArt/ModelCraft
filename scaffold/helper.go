package scaffold

import (
	"strings"
	"unicode"
)

const moduleName = "github.com/MarcelArt/app_standard"

// ToCamelCase converts a TitleCase string into titleCase
func ToCamelCase(input string) string {
	if input == "" {
		return ""
	}

	var titleCaseBuilder strings.Builder

	for i, r := range input {
		if unicode.IsUpper(r) {
			if i == 0 {
				titleCaseBuilder.WriteRune(unicode.ToLower(r))
			} else {
				titleCaseBuilder.WriteRune(r)
			}
		} else {
			titleCaseBuilder.WriteRune(r)
		}
	}

	return titleCaseBuilder.String()
}

// ToSeparateByCharLowered converts a TitleCase string into title_case
func ToSeparateByCharLowered(input string, char rune) string {
	if input == "" {
		return ""
	}
	if char == 0 {
		char = '_'
	}

	var titleCaseUnderscoreBuilder strings.Builder

	for i, r := range input {
		if unicode.IsUpper(r) {
			if i != 0 {
				titleCaseUnderscoreBuilder.WriteRune(char)
			}
			titleCaseUnderscoreBuilder.WriteRune(unicode.ToLower(r))
		} else {
			titleCaseUnderscoreBuilder.WriteRune(r)
		}
	}

	return titleCaseUnderscoreBuilder.String()
}

// PluralizeWord converts a singular word to its plural form
func PluralizeWord(word string) string {
	if word == "" {
		return ""
	}

	// Simple pluralization rules
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "sh") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") {
		return word + "es"
	} else if strings.HasSuffix(word, "y") && len(word) > 1 && !isVowel(rune(word[len(word)-2])) {
		return word[:len(word)-1] + "ies"
	} else {
		return word + "s"
	}
}

// Helper function to check if a rune is a vowel
func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}
