//  conversão: multiplicar Celsius por 1,8 (ou 9/5) e depois somar 32

package main

import "fmt"

func menu() {
	fmt.Println("=== CONVERTOR DE TEMPERATURA ===")
	fmt.Println("1 - converter")
	fmt.Println("2 - sair")
	fmt.Println("================================")
	fmt.Print("opcao escolhida: ")
}

func separador() {
	fmt.Println("================================")
}

func converter() {
	fmt.Print("quantos graus quer converter para Fahrenheit: ")
	var celcius float64
	fmt.Scanln(&celcius)

	resultado := (celcius * 1.8) + 32

	fmt.Println("resultado: ", resultado)
}

func main() {
	var opcao int
	for {
		menu()
		fmt.Scanln(&opcao)

		switch opcao {
		case 0:
			fmt.Println("saindo...")
			separador()
			return
		case 1:
			converter()
		}
	}
}
