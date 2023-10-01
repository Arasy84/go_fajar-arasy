package calculate

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(10, 5)
	if result != 15 {
		t.Errorf("Expected: 15, but got: %f", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10, 5)
	if result != 5 {
		t.Errorf("Expected: 5, but got: %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(10, 5)
	if result != 50 {
		t.Errorf("Expected: 50, but got: %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(10, 5)
	if err != nil || result != 2 {
		t.Errorf("Expected: 2 without error, but got: %.2f with error: %v", result, err)
	}

	_, err = Division(10, 0)
	if err == nil {
		t.Errorf("Expected error when dividing by zero, but no error found")
	}
}
