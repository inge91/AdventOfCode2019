package main

import "testing"

func assertEqual(i, j int, t *testing.T) {
	if i != j {
		t.Errorf("error")
	}
}

func TestGetRequiredFull(t *testing.T) {
	assertEqual(GetRequiredFull(12), 2, t)
	assertEqual(GetRequiredFull(14), 2, t)
	assertEqual(GetRequiredFull(1969), 654, t)
	assertEqual(GetRequiredFull(100756), 33583, t)
}

func TestGetRecursiveRequiredFull(t *testing.T) {
	assertEqual(GetRecursiveRequiredFull(12), 2, t)
	assertEqual(GetRecursiveRequiredFull(1969), 966, t)
	assertEqual(GetRecursiveRequiredFull(100756), 50346, t)
}
