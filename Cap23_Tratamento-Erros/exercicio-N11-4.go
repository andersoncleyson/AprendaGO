package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	long string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main(){
	_, err := sqrt(-10.34)
	if err != nil {
		log.Println(err)
	}
	

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erroNovo := fmt.Errorf("Deu erro no valor: %v", f)
		return 0, sqrtError{"", "", erroNovo}
	}

	return 42, nil
}