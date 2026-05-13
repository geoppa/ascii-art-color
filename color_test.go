package main

import "testing"

func TestAnsiStyle_RGB(t *testing.T) {
	s, err := ansiStyle("rgb(255, 0, 0)")
	if err != nil {
		t.Fatal(err)
	}
	if s != "38;2;255;0;0" {
		t.Fatalf("got %q", s)
	}
}

func TestAnsiStyle_HSL_Red(t *testing.T) {
	s, err := ansiStyle("hsl(0, 100%, 50%)")
	if err != nil {
		t.Fatal(err)
	}
	if s != "38;2;255;0;0" {
		t.Fatalf("got %q want 38;2;255;0;0", s)
	}
}

func TestAnsiStyle_Hex(t *testing.T) {
	s, err := ansiStyle("#00FF00")
	if err != nil {
		t.Fatal(err)
	}
	if s != "38;2;0;255;0" {
		t.Fatalf("got %q", s)
	}
}
