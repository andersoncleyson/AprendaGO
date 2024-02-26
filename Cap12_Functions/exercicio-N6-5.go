package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * 2
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

func info(f figura) {
	fmt.Println(f.area())
}

func main() {
	square := quadrado{3}
	circle := circulo{4}

	info(square)
	info(circle)
}
