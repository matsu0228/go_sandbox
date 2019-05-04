package main

import "testing"

func TestZigZag(t *testing.T) {

	s := "PAYPALISHIRING"

	// 4
	want := "PINALSIGYAHRPI"
	got := convert(s, 4)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}

	// 3
	want = "PAHNAPLSIIGYIR"
	got = convert(s, 3)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}

	// short
	s = "A"
	want = "A"
	got = convert(s, 1)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}

	// short
	s = "AB"
	want = "AB"
	got = convert(s, 1)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
	s = "AB"
	want = "AB"
	got = convert(s, 2)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
	s = "ABC"
	want = "ACB"
	got = convert(s, 2)
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
}
