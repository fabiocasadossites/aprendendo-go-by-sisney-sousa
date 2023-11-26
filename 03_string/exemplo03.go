package main

import "fmt"

func main()  {
	frase := "Ana comeu a banana"
	tamanho := len(frase)
	fmt.Printf("Tamanho da frase: %d\n",tamanho)
	
	primeiraPalavra := frase[0:3]
	fmt.Printf("Primeira palavra: %s\n",primeiraPalavra)

	segundaPalavra := frase[4:9]
	fmt.Printf("Segunda palavra: %s\n",segundaPalavra)

	terceiraPalavra := frase[12:]
	fmt.Printf("Terceira palavra: %s\n",terceiraPalavra)
}