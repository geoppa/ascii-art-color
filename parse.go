package main

import (
	"fmt"
	"strings"
)

const usageMsg = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
EX: go run . --color=<c1> <sub1> --color=<c2> <sub2> "text"
`

// ColorRule is one --color=<color> with optional substring (empty = whole string).
type ColorRule struct {
	Color     string
	Substring string
}

// Config holds parsed CLI state for one run.
type Config struct {
	Text   string
	Banner string
	Rules  []ColorRule // empty = no coloring
}

func isBannerName(s string) bool {
	switch strings.ToLower(s) {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}

// ParseConfig interprets argv like os.Args (including program name at [0]).
func ParseConfig(argv []string) (Config, error) {
	var c Config
	if len(argv) < 2 {
		return c, fmt.Errorf("usage")
	}
	if len(argv) == 2 && strings.HasPrefix(argv[1], "--color=") {
		return c, fmt.Errorf("usage")
	}
	last := argv[len(argv)-1]
	middle := argv[1 : len(argv)-1]

	i := 0
	if i < len(middle) && isBannerName(middle[i]) {
		c.Banner = strings.ToLower(middle[i]) + ".txt"
		i++
	} else {
		c.Banner = "standard.txt"
	}

	if i >= len(middle) {
		c.Text = last
		return c, nil
	}

	for i < len(middle) {
		arg := middle[i]
		if !strings.HasPrefix(arg, "--color=") {
			return c, fmt.Errorf("usage")
		}
		colorPart := strings.TrimPrefix(arg, "--color=")
		if colorPart == "" {
			return c, fmt.Errorf("usage")
		}
		i++

		var sub string
		if i < len(middle) && !strings.HasPrefix(middle[i], "--color=") {
			if strings.HasPrefix(middle[i], "--") {
				return c, fmt.Errorf("usage")
			}
			sub = middle[i]
			i++
		}

		c.Rules = append(c.Rules, ColorRule{Color: colorPart, Substring: sub})
	}

	c.Text = last
	return c, nil
}
