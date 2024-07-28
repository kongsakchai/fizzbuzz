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

func TestCase2(t *testing.T) {
	// Arrange
	give := 2
	want := "2"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase3(t *testing.T) {
	// Arrange
	give := 3
	want := "Fizz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase4(t *testing.T) {
	// Arrange
	give := 4
	want := "4"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase5(t *testing.T) {
	// Arrange
	give := 5
	want := "Buzz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase6(t *testing.T) {
	// Arrange
	give := 6
	want := "Fizz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase7(t *testing.T) {
	// Arrange
	give := 7
	want := "7"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}
