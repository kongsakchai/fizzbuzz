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

func TestCase8(t *testing.T) {
	// Arrange
	give := 8
	want := "8"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase9(t *testing.T) {
	// Arrange
	give := 9
	want := "Fizz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase10(t *testing.T) {
	// Arrange
	give := 10
	want := "Buzz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase11(t *testing.T) {
	// Arrange
	give := 11
	want := "11"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase12(t *testing.T) {
	// Arrange
	give := 12
	want := "Fizz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase13(t *testing.T) {
	// Arrange
	give := 13
	want := "13"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase14(t *testing.T) {
	// Arrange
	give := 14
	want := "14"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase15(t *testing.T) {
	// Arrange
	give := 15
	want := "FizzBuzz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}

func TestCase18(t *testing.T) {
	// Arrange
	give := 18
	want := "Fizz"

	// Act
	got := FizzBuzz(give)

	// Assert
	if got != want {
		t.Errorf("give %d expected %s, actual %s", give, want, got)
	}
}
