package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}
func exibeIntroducao() {
	fmt.Println("Olá Mundo")
	name := "Doux"
	var versao float32 = 1.1
	var tipo = 1.001

	fmt.Println("O meu nome é: ", name)
	fmt.Println("Este programa está na versao: ", versao)
	fmt.Println("O tipo da variável é: ", reflect.TypeOf(tipo))

	fmt.Println("\n \n ")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("\n  \n ")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)

	return comandoLido
}
