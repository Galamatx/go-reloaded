package transforms

import "strings"

func ApplyArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]

		startIdx := 0
		for startIdx < len(word) && (word[startIdx] == '\'' || word[startIdx] == '"' || word[startIdx] == '(') {
			startIdx++
		}

		cWord := word[startIdx:]
		if cWord != "a" && cWord != "A" && cWord != "an" && cWord != "An" {
			continue
		}

		nextIdx := i + 1
		for nextIdx < len(words) && words[nextIdx] == "" {
			nextIdx++
		}
		if nextIdx >= len(words) {
			continue
		}

		nextWord := strings.ToLower(words[nextIdx])
		for len(nextWord) > 0 && (nextWord[0] == '\'' || nextWord[0] == '"' || nextWord[0] == '(') {
			nextWord = nextWord[1:]
		}
		if len(nextWord) == 0 {
			continue
		}

		var isVowelOrH bool
		isVowelOrH = isVowel(nextWord[0]) || nextWord[0] == 'h'

		prefix := word[:startIdx]
		switch {
		case (cWord == "a" || cWord == "A") && isVowelOrH:
			if cWord == "a" {
				words[i] = prefix + "an"
			} else {
				words[i] = prefix + "An"
			}
		case (cWord == "an" || cWord == "An") && !isVowelOrH:
			if cWord == "an" {
				words[i] = prefix + "a"
			} else {
				words[i] = prefix + "A"
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
