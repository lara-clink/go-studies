package main

import "fmt"

func main() {
	nome := "Lara"
	versao := 1.1
	fmt.Println("ola senhora,", nome)
	fmt.Println("Este progama esta na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)
}
