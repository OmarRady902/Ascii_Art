package main

import (
	"ASCII_ART/banner"
	"ASCII_ART/render"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage message : go run . <text> <banner> ")
		os.Exit(1)
	}

	text := os.Args[1]
	if text == "" {
		return
	}

	bannerFile := "standard"

	if len(os.Args) == 3 {
		bannerFile = os.Args[2]
	}

	content, err := banner.LoadBannerFile("banners/" + bannerFile + ".txt")
	if err != nil {
		fmt.Println("There is a problem in reading file:\nThe Problem\n", err)
		os.Exit(1)
	}

	charMap := banner.BuildCharacterMap(content)

	text = strings.ReplaceAll(text, `\n`, "\n")
	if strings.Trim(text, "\n") == "" {
		count := strings.Count(text, "\n")
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return
	}
	for _, ch := range text {
		if ch != '\n' && (ch < 32 || ch > 126) {
			fmt.Println("Error: unsupported character")
			os.Exit(1)
		}
	}

	parts := strings.Split(text, "\n")

	for _, part := range parts {
		if part == "" {
			fmt.Println()
		} else {
			result := render.RenderText(part, charMap)
			fmt.Print(result)
		}
	}
}
