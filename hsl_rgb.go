package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// parseRGBFunc parses "rgb(255, 0, 0)" (case-insensitive, flexible spaces).
func parseRGBFunc(s string) (r, g, b int, err error) {
	s = strings.TrimSpace(strings.ToLower(s))
	if !strings.HasPrefix(s, "rgb(") || !strings.HasSuffix(s, ")") {
		return 0, 0, 0, fmt.Errorf("not rgb()")
	}
	inner := strings.TrimSpace(s[4 : len(s)-1])
	parts := strings.Split(inner, ",")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("rgb needs 3 components")
	}
	var vals [3]int
	for i := 0; i < 3; i++ {
		v, err := strconv.Atoi(strings.TrimSpace(parts[i]))
		if err != nil || v < 0 || v > 255 {
			return 0, 0, 0, fmt.Errorf("invalid rgb component")
		}
		vals[i] = v
	}
	return vals[0], vals[1], vals[2], nil
}

// parseHSLFunc parses "hsl(0, 100%, 50%)" (degrees and two percentages).
func parseHSLFunc(s string) (r, g, b int, err error) {
	s = strings.TrimSpace(strings.ToLower(s))
	if !strings.HasPrefix(s, "hsl(") || !strings.HasSuffix(s, ")") {
		return 0, 0, 0, fmt.Errorf("not hsl()")
	}
	inner := strings.TrimSpace(s[4 : len(s)-1])
	parts := strings.Split(inner, ",")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("hsl needs 3 components")
	}
	hStr := strings.TrimSpace(parts[0])
	sStr := strings.Trim(strings.TrimSpace(parts[1]), "%")
	lStr := strings.Trim(strings.TrimSpace(parts[2]), "%")

	hf, err := strconv.ParseFloat(hStr, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	sf, err := strconv.ParseFloat(sStr, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	lf, err := strconv.ParseFloat(lStr, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	if sf < 0 || sf > 100 || lf < 0 || lf > 100 {
		return 0, 0, 0, fmt.Errorf("hsl s/l out of range")
	}
	r, g, b = hslToRGB(hf, sf/100, lf/100)
	return r, g, b, nil
}

func hslToRGB(h, s, l float64) (int, int, int) {
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r1, g1, b1 float64
	switch {
	case h < 60:
		r1, g1, b1 = c, x, 0
	case h < 120:
		r1, g1, b1 = x, c, 0
	case h < 180:
		r1, g1, b1 = 0, c, x
	case h < 240:
		r1, g1, b1 = 0, x, c
	case h < 300:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}
	return clamp255((r1 + m) * 255), clamp255((g1 + m) * 255), clamp255((b1 + m) * 255)
}

func clamp255(v float64) int {
	v = math.Round(v)
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return int(v)
}
