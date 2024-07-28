package main

import "testing"

func TestCase1(t *testing.T) {
	// Arrange
	give := 1
	want := "1"

	// Act
	got := FizzBuzz(1)

	// Assert
	if want != got {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}

}
