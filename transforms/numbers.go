package transforms

import (
	"strconv"
	"strings"
)

func ApplyNumbers(text string) string {
	text = strings.ReplaceAll(text, "(hex)", " (hex) ")
	text = strings.ReplaceAll(text, "(bin)", " (bin) ")

	words := strings.Fields(text)

	for i, word := range words {
		switch word {
		case "(hex)":
			prevIndex := i - 1
			for prevIndex >= 0 && words[prevIndex] == "" {
				prevIndex--
			}
			if prevIndex >= 0 {
				num, punct := splitTrailingPunct(words[prevIndex])
				if isHexStr(num) {
					words[prevIndex] = strconv.Itoa(hexTOdec(num)) + punct
				}
				words[i] = ""
			}
		case "(bin)":
			prevIndex := i - 1
			for prevIndex >= 0 && words[prevIndex] == "" {
				prevIndex--
			}
			if prevIndex >= 0 {
				num, punct := splitTrailingPunct(words[prevIndex])
				if isBinStr(num) {
					words[prevIndex] = strconv.Itoa(binTOdec(num)) + punct
				}
				words[i] = ""
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

func splitTrailingPunct(s string) (string, string) {
	i := len(s) - 1
	for i >= 0 && (s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';') {
		i--
	}
	return s[:i+1], s[i+1:]
}

func isHexStr(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F') {
			return false
		}
	}
	return true
}

func isBinStr(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return false
		}
	}
	return true
}

func hexTOdec(hex string) int {
	result := 0

	if len(hex) >= 2 && hex[0] == '0' && (hex[1] == 'x' || hex[1] == 'X') {
		hex = hex[2:]
	}

	for i := 0; i < len(hex); i++ {
		char := hex[i]
		var val int

		if char >= '0' && char <= '9' {
			val = int(char - '0')
		} else if char >= 'a' && char <= 'f' {
			val = int(char - 'a' + 10)
		} else if char >= 'A' && char <= 'F' {
			val = int(char - 'A' + 10)
		} else {
			return 0
		}

		result = (result * 16) + val
	}

	return result
}

func binTOdec(bin string) int {
	result := 0

	for i := 0; i < len(bin); i++ {
		val := int(bin[i] - '0')
		result = (result * 2) + val
	}
	return result
}
