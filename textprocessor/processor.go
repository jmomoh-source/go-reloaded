package textprocessor

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Process applies all text transformations to the input string.
func Process(text string) string {
	words := tokenize(text)
	words = applyFlags(words)
	words = fixPunctuation(words)
	words = fixQuotes(words)
	words = fixArticles(words)
	return strings.Join(words, " ")
}

// tokenize splits text into tokens by whitespace.
func tokenize(text string) []string {
	return strings.Fields(text)
}

// applyFlags scans for (hex), (bin), (up), (low), (cap) and their (xx, N) variants.
func applyFlags(words []string) []string {
	// We use a regex to match flag patterns like (hex), (bin), (up), (low), (cap), (up, 3), etc.
	singleFlag := regexp.MustCompile(`^\((hex|bin|up|low|cap)\)$`)
	numberedFlag := regexp.MustCompile(`^\((up|low|cap),\s*(\d+)\)$`)

	// We may also encounter a flag split across tokens like "(cap," and "6)"
	// due to whitespace tokenization. We need to handle multi-token flags.
	// Pattern: "(up," or "(low," or "(cap," followed by a token like "6)"
	flagStart := regexp.MustCompile(`^\((up|low|cap),$`)
	flagEnd := regexp.MustCompile(`^(\d+)\)$`)

	var result []string

	i := 0
	for i < len(words) {
		word := words[i]

		// Check for single-token flag: (hex), (bin), (up), (low), (cap)
		if m := singleFlag.FindStringSubmatch(word); m != nil {
			flag := m[1]
			applyTransformation(&result, flag, 1)
			i++
			continue
		}

		// Check for single-token numbered flag: (up,3) or (cap, 6)
		if m := numberedFlag.FindStringSubmatch(word); m != nil {
			flag := m[1]
			n, _ := strconv.Atoi(m[2])
			applyTransformation(&result, flag, n)
			i++
			continue
		}

		// Check for multi-token numbered flag: "(cap," followed by "6)"
		if m := flagStart.FindStringSubmatch(word); m != nil && i+1 < len(words) {
			nextWord := words[i+1]
			if mn := flagEnd.FindStringSubmatch(nextWord); mn != nil {
				flag := m[1]
				n, _ := strconv.Atoi(mn[1])
				applyTransformation(&result, flag, n)
				i += 2
				continue
			}
		}

		result = append(result, word)
		i++
	}

	return result
}

// applyTransformation applies a flag transformation to the last N words in result.
func applyTransformation(result *[]string, flag string, n int) {
	r := *result
	count := n
	if count > len(r) {
		count = len(r)
	}

	startIdx := len(r) - count
	for j := startIdx; j < len(r); j++ {
		switch flag {
		case "hex":
			val, err := strconv.ParseInt(r[j], 16, 64)
			if err == nil {
				r[j] = fmt.Sprintf("%d", val)
			}
		case "bin":
			val, err := strconv.ParseInt(r[j], 2, 64)
			if err == nil {
				r[j] = fmt.Sprintf("%d", val)
			}
		case "up":
			r[j] = strings.ToUpper(r[j])
		case "low":
			r[j] = strings.ToLower(r[j])
		case "cap":
			r[j] = capitalize(r[j])
		}
	}
	*result = r
}

// capitalize returns the word with the first letter uppercased and the rest lowercased.
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(strings.ToLower(word))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// isPunctuation checks if a string is a punctuation mark.
func isPunctuation(s string) bool {
	for _, r := range s {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			continue
		default:
			return false
		}
	}
	return len(s) > 0
}

// fixPunctuation ensures punctuation is attached to the previous word.
// Groups of punctuation like "..." or "!?" are kept together and attached to previous word.
// Also handles punctuation glued to the start of a word (e.g., ",what" -> attach "," to prev word, keep "what").
func fixPunctuation(words []string) []string {
	if len(words) == 0 {
		return words
	}

	// First pass: split tokens that have leading punctuation glued to a word
	// e.g., ",what" -> ",", "what"
	var expanded []string
	for _, word := range words {
		leading := ""
		rest := word
		for len(rest) > 0 && isPunctuationChar(rune(rest[0])) {
			leading += string(rest[0])
			rest = rest[1:]
		}
		if leading != "" && rest != "" {
			expanded = append(expanded, leading, rest)
		} else {
			expanded = append(expanded, word)
		}
	}

	// Second pass: attach punctuation-only tokens to the previous word
	var result []string
	for _, word := range expanded {
		if isPunctuation(word) && len(result) > 0 {
			result[len(result)-1] += word
		} else {
			result = append(result, word)
		}
	}

	return result
}

// isPunctuationChar checks if a single rune is a punctuation character.
func isPunctuationChar(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';':
		return true
	}
	return false
}

// fixQuotes handles single-quote pairing: ' word ' -> 'word'
// and ' multiple words ' -> 'multiple words'
func fixQuotes(words []string) []string {
	if len(words) == 0 {
		return words
	}

	var result []string

	i := 0
	for i < len(words) {
		if words[i] == "'" {
			// Find the matching closing quote
			closeIdx := -1
			for j := i + 1; j < len(words); j++ {
				if words[j] == "'" {
					closeIdx = j
					break
				}
			}

			if closeIdx != -1 && closeIdx > i+1 {
				// Collect words between quotes
				inner := words[i+1 : closeIdx]
				quoted := "'" + strings.Join(inner, " ") + "'"
				result = append(result, quoted)
				i = closeIdx + 1
			} else if closeIdx == i+1 {
				// Empty quotes ''
				result = append(result, "''")
				i = closeIdx + 1
			} else {
				// No matching close quote found, keep as-is
				result = append(result, words[i])
				i++
			}
		} else {
			result = append(result, words[i])
			i++
		}
	}

	return result
}

// fixArticles converts "a" to "an" when the next word starts with a vowel or 'h'.
func fixArticles(words []string) []string {
	if len(words) == 0 {
		return words
	}

	for i := 0; i < len(words)-1; i++ {
		lower := strings.ToLower(words[i])
		// Only match standalone "a" (not "a" that's part of another word attached with punctuation)
		if lower != "a" {
			continue
		}

		nextWord := words[i+1]
		// Strip any leading quote marks to get the actual first letter
		cleaned := strings.TrimLeft(nextWord, "'\"")
		if len(cleaned) == 0 {
			continue
		}

		firstChar := unicode.ToLower(rune(cleaned[0]))
		if isVowelOrH(firstChar) {
			// Preserve original case
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}

	return words
}

// isVowelOrH checks if a rune is a vowel or 'h'.
func isVowelOrH(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	}
	return false
}
