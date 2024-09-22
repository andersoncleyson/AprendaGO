package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(1, 2, 3, 4, 5)
	resultado := 15

	if teste != 15 {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
