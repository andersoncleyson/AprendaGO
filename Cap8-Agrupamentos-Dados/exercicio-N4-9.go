package main

import "fmt"

func main(){
	pessoas := map[string][]string {
		"Elliot Alderson": {
			"Hacker", "escrever diário",
		},
		"Shayla Nico": {
			"Vendar arte na praia", "cantar",
		},
		"Leon": {
			"Fazer música", "trabalhar pra dark army",
		},
	}

	pessoas["Angela Moss"] = []string{"Auto ajuda", "correr"}

	for key, value := range pessoas {
		fmt.Printf("%s:\n", key)
		for _, h := range value {
			fmt.Printf("\t- %s\n", h)
		}
		fmt.Printf("\n")
	}

}