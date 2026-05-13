package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usageMsg)
		return
	}

	cfg, err := ParseConfig(os.Args)
	if err != nil {
		fmt.Print(usageMsg)
		return
	}

	input := cfg.Text
	if input == "" {
		return
	}

	lines, err := readBanner(cfg.Banner)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	parts := strings.Split(input, "\\n")

	onlyNewlines := true
	for _, part := range parts {
		if part != "" {
			onlyNewlines = false
			break
		}
	}
	if onlyNewlines {
		for i := 0; i < len(parts)-1; i++ {
			fmt.Println()
		}
		return
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}

		var opens []string
		if len(cfg.Rules) > 0 {
			opens, err = perRuneANSIOpens(part, cfg.Rules)
			if err != nil {
				fmt.Print(usageMsg)
				return
			}
		}

		runes := []rune(part)
		for i := 0; i < 8; i++ {
			for j, char := range runes {
				startLine := int(char-32)*9 + 1

				var seg string
				if startLine >= 1 && startLine+i < len(lines) {
					seg = lines[startLine+i]
				}

				if len(opens) > 0 && j < len(opens) && opens[j] != "" {
					fmt.Print(opens[j] + seg + ansiReset)
				} else {
					fmt.Print(seg)
				}
			}
			fmt.Println()
		}
	}
}

func readBanner(bannerName string) ([]string, error) {
	file, err := os.Open(bannerName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
