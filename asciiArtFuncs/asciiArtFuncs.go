package asciiArtWeb

import (
	"os"
	"strings"
)

var asciiMap = map[int]int {	}
var asciiLines []string

func ReadBanner(bannerName string) string {

	for lineNr, charCode := 1, 0; charCode <= 126; lineNr, charCode = lineNr + 9, charCode + 1 {
		asciiMap[32+charCode] = lineNr
	}

	if bannerName == "thinkertoy" {
		file, err := os.ReadFile("thinkertoy.txt")
		if err != nil {
			return string("Error: " + err.Error())
		}
		thinkertoy := string(file)
		asciiLines = strings.Split(thinkertoy, "\n")
	} else if bannerName == "shadow" {
		file, err := os.ReadFile("shadow.txt")
		if err != nil {
			return string("Error: " + err.Error())
		}
		shadow := string(file)
		asciiLines = strings.Split(shadow, "\n")
	} else {
		file, err := os.ReadFile("standard.txt")
		if err != nil {
			return string("Error: " + err.Error())
		}
		standard := string(file)
		asciiLines = strings.Split(standard, "\n")
	}
	return ""
}

func FormatTextToAscii(inputText string) string {
	lines := strings.Split(inputText, "\n")
	formatedToAscii := ""
	for _, line := range lines {
		for i := 0; i < 8; i++ {
			for _, char := range line {
				mapPos := asciiMap[int(char)]
				formatedToAscii += asciiLines[mapPos+i]
			}
			formatedToAscii += "\n"
		}
	}
	return formatedToAscii
}
