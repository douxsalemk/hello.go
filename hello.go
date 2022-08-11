package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Olá Mundo")
	name := "Doux"
	var versao float32 = 1.1
	var tipo = 1.001

	fmt.Println("O meu nome é: ", name)
	fmt.Println("Este programa está na versao: ", versao)
	fmt.Println("O tipo da variável é: ", reflect.TypeOf(tipo))

	fmt.Println("\n \n ")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	fmt.Println("\n  \n ")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereco da minha variavel comando é: ", &comando)
	fmt.Println("O comando escolhido foi: ", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}

}
