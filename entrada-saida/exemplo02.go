package main

import "fmt"

func main()  {
	var nome string

	fmt.Println("Digite o seu nome:")
	fmt.Scanf("%s", &nome)
	fmt.Printf("Seja bem-vindo(a), %s!\n", nome)

	//var serie string
	serie := "Lógica de programação com Go"
	fmt.Printf("Seja bem-vindo(a) `a série %s, %s!\n", serie, nome)
	serie = "Corte e costura"
}