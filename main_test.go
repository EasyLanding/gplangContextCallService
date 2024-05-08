package main

import "testing"

func TestHelloWorld(t *testing.T) {
	execting := `Hello World`

	result := HelloWorld()
	if result != execting {
		t.Errorf("HelloWorld() returned %s, expected %s", result, execting)
	}
}
