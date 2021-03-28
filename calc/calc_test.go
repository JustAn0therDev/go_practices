package calc

import (
	"testing"
)

func SumTest(t *testing.T) {
	result, err := sum(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a sum value but got %v", err)
	}
}

func SubTest(t *testing.T) {
	result, err := sub(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a subtracted value but got %v", err)
	}
}

func MultTest(t *testing.T) {
	result, err := mult(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a product value but got %v", err)
	}
}

func DivTest(t *testing.T) {
	result, err := div(1, 2)
	if result == 0 || err != nil {
		t.Fatalf("expected a division value but got %v", err)
	}
}