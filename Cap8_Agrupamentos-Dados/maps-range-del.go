package main

import "fmt"

func main(){
	qualquercoisa := map[int]string{
		123: "muito legal",
		902: "menos legal um pouco",
		324: "esse é massa",
		18: "idade de ir pra gesta",
	}

	fmt.Println(qualquercoisa)

	for i, v := range qualquercoisa {
		fmt.Println(i, v)
	}

	delete(qualquercoisa, 324)

	fmt.Println(qualquercoisa)
}