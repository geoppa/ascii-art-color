package main

import (
	"strings"
	"testing"
)

func TestParseConfig_SingleString(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "hello"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "hello" || cfg.Banner != "standard.txt" || len(cfg.Rules) != 0 {
		t.Fatalf("%+v", cfg)
	}
}

func TestParseConfig_BannerAndString(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "thinkertoy", "hi"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "hi" || cfg.Banner != "thinkertoy.txt" || len(cfg.Rules) != 0 {
		t.Fatalf("%+v", cfg)
	}
}

func TestParseConfig_ColorWholeString(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "--color=red", "abc"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "abc" || len(cfg.Rules) != 1 || cfg.Rules[0].Color != "red" || cfg.Rules[0].Substring != "" {
		t.Fatalf("%+v", cfg)
	}
}

func TestParseConfig_ColorSubstring(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "--color=red", "kit", "a king kitten have kit"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "a king kitten have kit" || len(cfg.Rules) != 1 || cfg.Rules[0].Color != "red" || cfg.Rules[0].Substring != "kit" {
		t.Fatalf("%+v", cfg)
	}
}

func TestParseConfig_BannerColorSubstring(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "shadow", "--color=green", "o", "go"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Banner != "shadow.txt" || cfg.Text != "go" || len(cfg.Rules) != 1 || cfg.Rules[0].Substring != "o" || cfg.Rules[0].Color != "green" {
		t.Fatalf("%+v", cfg)
	}
}

func TestParseConfig_TwoColors(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "--color=red", "a", "--color=blue", "b", "ab"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "ab" || len(cfg.Rules) != 2 {
		t.Fatalf("%+v", cfg)
	}
	if cfg.Rules[0].Color != "red" || cfg.Rules[0].Substring != "a" || cfg.Rules[1].Color != "blue" || cfg.Rules[1].Substring != "b" {
		t.Fatalf("%+v", cfg.Rules)
	}
}

func TestParseConfig_TwoColorsWholeThenSubstring(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", "--color=red", "--color=blue", "x", "ax"})
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Rules) != 2 {
		t.Fatalf("%+v", cfg.Rules)
	}
	if cfg.Rules[0].Color != "red" || cfg.Rules[0].Substring != "" {
		t.Fatalf("rule0 %+v", cfg.Rules[0])
	}
	if cfg.Rules[1].Color != "blue" || cfg.Rules[1].Substring != "x" {
		t.Fatalf("rule1 %+v", cfg.Rules[1])
	}
}

func TestParseConfig_InvalidFlag(t *testing.T) {
	_, err := ParseConfig([]string{"ascii-art", "--colour=red", "x"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseConfig_ColorWithoutEquals(t *testing.T) {
	_, err := ParseConfig([]string{"ascii-art", "--color", "red", "x"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseConfig_ExtraMiddleToken(t *testing.T) {
	_, err := ParseConfig([]string{"ascii-art", "--color=red", "a", "b", "c"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseConfig_LoneColorFlagTwoArgs(t *testing.T) {
	_, err := ParseConfig([]string{"ascii-art", "--color=red"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseConfig_EmptyStringArg(t *testing.T) {
	cfg, err := ParseConfig([]string{"ascii-art", ""})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Text != "" {
		t.Fatalf("got %q", cfg.Text)
	}
}

func TestColoredRunes_KittenKit(t *testing.T) {
	m := coloredRunes("a king kitten have kit", "kit")
	s := string([]rune("a king kitten have kit"))
	if len(m) != len([]rune(s)) {
		t.Fatalf("len %d vs runes %d", len(m), len([]rune(s)))
	}
	r := []rune("a king kitten have kit")
	idx := strings.Index(string(r), "kitten")
	if idx < 0 {
		t.Fatal("kitten not found")
	}
	runeIdx := len([]rune(string(r)[:idx]))
	for k := 0; k < 3; k++ {
		if !m[runeIdx+k] {
			t.Fatalf("expected kitten prefix colored at %d", runeIdx+k)
		}
	}
	last := strings.LastIndex(string(r), "kit")
	li := len([]rune(string(r)[:last]))
	for k := 0; k < 3; k++ {
		if !m[li+k] {
			t.Fatalf("expected trailing kit colored at %d", li+k)
		}
	}
}

func TestColoredRunes_EmptySubstringColorsAll(t *testing.T) {
	m := coloredRunes("abc", "")
	for i, v := range m {
		if !v {
			t.Fatalf("index %d not colored", i)
		}
	}
}

func TestPerRuneANSIOpens_LaterRuleOverlaps(t *testing.T) {
	rules := []ColorRule{
		{Color: "red", Substring: "ab"},
		{Color: "blue", Substring: "b"},
	}
	opens, err := perRuneANSIOpens("ab", rules)
	if err != nil {
		t.Fatal(err)
	}
	if len(opens) != 2 {
		t.Fatalf("len %d", len(opens))
	}
	// 'a' only red; 'b' matched by both -> last rule (blue) wins
	if !strings.Contains(opens[0], "31") {
		t.Fatalf("a should be red SGR, got %q", opens[0])
	}
	if !strings.Contains(opens[1], "34") {
		t.Fatalf("b should be blue SGR, got %q", opens[1])
	}
}
