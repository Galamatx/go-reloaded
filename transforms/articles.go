package transforms

import "strings"

func ApplyArticles(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		if (word == "a" || word == "A") && i+1 < len(words) {
			nextWord := words[i+1]
			if isVowel(nextWord[0]) || nextWord[0] == 'h' || nextWord[0] == 'H' {
				if word == "a" {
					words[i] = string("an")
				} else if word == "A" {
					words[i] = string("An")
				}
			}
		}
	}

	result := []string{}
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
	}
	return strings.Join(result, " ")
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'A' || c == 'e' || c == 'E' ||
		c == 'i' || c == 'I' || c == 'o' || c == 'O' ||
		c == 'u' || c == 'U' {
		return true
	} else {
		return false
	}
}
