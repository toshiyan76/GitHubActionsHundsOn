package main

import (
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "偶数" {
		t.Errorf("Expected even, but got %s", result)
	}
}
