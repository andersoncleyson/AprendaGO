package main

import "testing"

type test struct {
	data []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test {
		test{data: []int{1, 2, 3}, answer: 6},
		test{data: []int{10, 11, 12}, answer: 33},
		test{data: []int{-5, 0, 5, 10}, answer: 10},
	}
	for _, v := range tests {
		z := soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}



func TestSoma(t *testing.T) {
	teste := soma(1, 2, 3, 4, 5)
	resultado := 15

	if teste != 15 {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

