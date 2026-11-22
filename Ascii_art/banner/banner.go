package banner

import (
	"os"
	"strings"
)

func LoadBannerFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func BuildCharacterMap(content string) map[rune][]string {
	charMap := make(map[rune][]string, 95)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines)-1; i += 9 {
		char := rune(32 + i/9)
		charMap[char] = lines[i+1 : i+9]
	}

	return charMap
}
