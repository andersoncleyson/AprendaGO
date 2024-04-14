package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Esse é o programa do exercício de compilação cruzada.\nFoi compilado num Windows e agora está rodando num sistema:")
	fmt.Println(runtime.GOARCH, runtime.GOOS)
}
