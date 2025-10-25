package main

import "testing"

func TestSanitizeChirp(t *testing.T) {
	test_string := "This is a kerfuffle day, my man."
	expected := "This is a **** day, my man."
	output := sanitizeChirp(test_string)
	if output != expected {
		t.Fatalf("got %q, want %q", output, expected)
	}
}

func TestMultipleCurseWords(t *testing.T) {
	test_string := "This is a kerfuffle day, my man. I wanted to eat Sharbert but I saw Fornax."
	expected := "This is a **** day, my man. I wanted to eat **** but I saw Fornax."
	output := sanitizeChirp(test_string)
	if output != expected {
		t.Fatalf("got %q, want %q", output, expected)
	}
}

func TestPunctuation(t *testing.T) {
	test_string := "Oh Kerfuffle! I had a nice Sharbert with Fornax!"
	expected := "Oh Kerfuffle! I had a nice **** with Fornax!"
	output := sanitizeChirp(test_string)
	if output != expected {
		t.Fatalf("got %q, want %q", output, expected)
	}
}
