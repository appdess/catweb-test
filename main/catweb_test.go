package main

import (
	"testing"
)

func TestRandom(t *testing.T) {
	min := 0
	max := 10
	result := Random(min, max)
	if result < min || result > max {
		t.Errorf("Expected result between %d and %d, got %d", min, max, result)
	}
}
