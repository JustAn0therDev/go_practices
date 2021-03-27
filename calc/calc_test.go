package calc

import (
	"testing"
)

func SumTest(t *testing.T) {
	result, err := Sum(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a sum but got %v", err)
	}
}

func SubTest(t *testing.T) {
	result, err := Sub(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a subtracted value but got %v", err)
	}
}

func MultTest(t *testing.T) {
	result, err := Mult(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a product but got %v", err)
	}
}

func DivTest(t *testing.T) {
	result, err := Div(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a division but got %v", err)
	}
}