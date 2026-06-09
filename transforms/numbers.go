package transforms

import (
	"strconv"
	"strings"
)

func ApplyNumbers(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		switch word {
		case "(hex)":
			{
				if i > 0 {
					words[i-1] = strconv.Itoa(hexTOdec(words[i-1]))
					words[i] = ""
				}
			}
		case "(bin)":
			{
				if i > 0 {
					words[i-1] = strconv.Itoa(binTOdec(words[i-1]))
					words[i] = ""
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
