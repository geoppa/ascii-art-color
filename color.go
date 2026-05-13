package main

import (
	"fmt"
	"strconv"
	"strings"
)

const ansiReset = "\033[0m"

// ansiStyle returns SGR parameters for foreground (no CSI wrapper).
func ansiStyle(color string) (string, error) {
	c := strings.TrimSpace(strings.ToLower(color))
	if c == "" {
		return "", fmt.Errorf("empty color")
	}
	if strings.HasPrefix(c, "#") && len(c) == 7 {
		r, g, b, err := parseHexRGB(c[1:])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
	}
	if strings.HasPrefix(c, "rgb(") {
		r, g, b, err := parseRGBFunc(c)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
	}
	if strings.HasPrefix(c, "hsl(") {
		r, g, b, err := parseHSLFunc(c)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
	}
	switch c {
	case "red":
		return "31", nil
	case "green":
		return "32", nil
	case "yellow":
		return "33", nil
	case "blue":
		return "34", nil
	case "magenta":
		return "35", nil
	case "cyan":
		return "36", nil
	case "white":
		return "37", nil
	case "black":
		return "30", nil
	case "orange":
		// Truecolor approximation (named orange not in basic 8).
		return "38;2;255;165;0", nil
	case "purple":
		return "38;2;128;0;128", nil
	default:
		return "", fmt.Errorf("unknown color %q", color)
	}
}

func parseHexRGB(hex string) (r, g, b int, err error) {
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex length")
	}
	v, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	r = int(v >> 16 & 0xff)
	g = int(v >> 8 & 0xff)
	b = int(v & 0xff)
	return r, g, b, nil
}

func ansiOpen(color string) (string, error) {
	s, err := ansiStyle(color)
	if err != nil {
		return "", err
	}
	return "\033[" + s + "m", nil
}
