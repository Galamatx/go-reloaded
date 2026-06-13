package transforms

import "strings"

func ApplyPunctuation(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "" {
			continue
		}

		for len(word) > 0 && isPunctCh(word[0]) {
			if i > 0 {
				words[i-1] = words[i-1] + string(word[0])
			}
			word = word[1:]
		}
		words[i] = word

		if isPunct(words[i]) && i > 0 {
			words[i-1] = words[i-1] + words[i]
			words[i] = ""
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

func isPunct(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':' {
			return true
		}
	}

	return false
}

func isPunctCh(c byte) bool {
	if c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':' {
		return true
	}
	return false
}