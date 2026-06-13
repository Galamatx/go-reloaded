package transforms

import "strings"

func ApplyQuotes(text string) string {
	words := strings.Fields(text)

	openIndex := -1
	for i, word := range words {
		if word == "'" {
			if openIndex == -1 {
				openIndex = i
			} else {
				words[openIndex+1] = "'" + words[openIndex+1]
				words[i-1] = words[i-1] + "'"
				words[openIndex] = ""
				words[i] = ""
				openIndex = -1
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
