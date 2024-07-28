package main

import "testing"

func TestCase1(t *testing.T) {
	// Arrange
	give := 1
	want := "1"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}
