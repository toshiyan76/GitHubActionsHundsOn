package main

import (
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "偶数だよ" {
		t.Errorf("Expected even, but got %s", result)
	}
}
