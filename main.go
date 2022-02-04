package main

import (
	"fmt"
	"time"
)

func helloGoroutine(name string) {
	fmt.Println("Executing goroutine", name)
	time.Sleep(1 * time.Second)
}

func startGame() {
	fmt.Println("E começa o jogo...")
}

func systemPresentation() {
	fmt.Println("Simulador de partidas de Tenis")
	fmt.Println("-------------------------------")
	fmt.Println("Escolha uma das opções:")
	fmt.Println("1 - Jogar")
	fmt.Println("2 - Configurar")
	fmt.Println("-------------------------------")
	fmt.Println("0 - Finalizar Programa")
}

func setConfiguration() {
	fmt.Println("Configurado")
}

// Definição de variáveis
var SCOREPLAYER1 = 0
var SCOREPLAYER2 = 0
var GAME = 1
var SET = 1
var MATCH = 1

func main() {
	var opcao int = 10
	var controleMenu = true

	// Loop responsável por permancer no programa
	for controleMenu {
		systemPresentation()
		fmt.Scanln(&opcao)
		switch opcao {
		case 1:
			startGame()
		case 2:
			setConfiguration()
		case 0:
			controleMenu = false
		}

		go helloGoroutine("G1")
		go helloGoroutine("G2")
		time.Sleep(1e9)
	}

}
