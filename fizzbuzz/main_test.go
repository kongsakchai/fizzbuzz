package main

import "testing"

func Test(t *testing.T) {
	// Arrange
	give := 1
	want := "1"

	// Act
	got := FizzBuzz(give)

	// Assert
	if want != got {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}
