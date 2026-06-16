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

		consonantSoundWords := map[string]bool{
			"user": true, "university": true, "unique": true, "union": true,
			"unit": true, "uniform": true, "united": true, "universal": true,
			"universe": true, "url": true, "european": true, "euro": true,
			"one": true, "once": true, "and": true, "or": true, "but": true,
		}

		vowelSoundWords := map[string]bool{
			"hour": true, "honest": true, "honour": true, "honor": true,
			"heir": true, "herb": true, "homage": true,
		}

		var isVowelOrH bool
		if consonantSoundWords[nextWord] {
			isVowelOrH = false
		} else if vowelSoundWords[nextWord] {
			isVowelOrH = true
		} else {
			isVowelOrH = isVowel(nextWord[0]) || nextWord[0] == 'h'
		}

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
