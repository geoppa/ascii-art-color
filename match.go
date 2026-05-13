package main

// coloredRunes marks rune indices in s that belong to non-overlapping occurrences of sub.
// If sub is empty, every index is marked (whole-string coloring).
func coloredRunes(s, sub string) []bool {
	sr := []rune(s)
	out := make([]bool, len(sr))
	if sub == "" {
		for i := range out {
			out[i] = true
		}
		return out
	}
	subr := []rune(sub)
	if len(subr) == 0 || len(subr) > len(sr) {
		return out
	}
	for i := 0; i <= len(sr)-len(subr); {
		match := true
		for j := 0; j < len(subr); j++ {
			if sr[i+j] != subr[j] {
				match = false
				break
			}
		}
		if match {
			for j := 0; j < len(subr); j++ {
				out[i+j] = true
			}
			i += len(subr)
			continue
		}
		i++
	}
	return out
}

// perRuneANSIOpens returns one CSI open sequence per rune of part (empty = no color).
// Rules are applied in order; later rules overwrite overlapping runes.
func perRuneANSIOpens(part string, rules []ColorRule) ([]string, error) {
	if len(rules) == 0 {
		return nil, nil
	}
	runes := []rune(part)
	out := make([]string, len(runes))
	for _, rule := range rules {
		open, err := ansiOpen(rule.Color)
		if err != nil {
			return nil, err
		}
		mark := coloredRunes(part, rule.Substring)
		for j := range runes {
			if j < len(mark) && mark[j] {
				out[j] = open
			}
		}
	}
	return out, nil
}
