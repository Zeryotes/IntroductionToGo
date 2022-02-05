package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Goroutine inicializadora
func initializer(state chan int) {
	fmt.Println("Iniciando a partida!")
	state <- 0
}

// Goroutine do jogador1
func player1(state chan int) {
	fmt.Println("----------- PLAYER 1 -----------")
	rand.Seed(time.Now().UnixNano() + 5)
	stateNow := <-state

	if stateNow == 0 {
		fmt.Println("Player1 joga a bola")
	} else {
		if isScore(stateNow) {
			SCOREPLAYER1++
			fmt.Printf("PLAYER 1 SCORED (%d)\n", SCOREPLAYER1)
		}
	}
	chanceOnPlay := rand.Intn(1000) + 1
	fmt.Println("Chance da jogada:", chanceOnPlay)
	state <- chanceOnPlay
	time.Sleep(1)
}

// Goroutine do jogador2
func player2(state chan int) {
	fmt.Println("----------- PLAYER 2 -----------")
	rand.Seed(time.Now().UnixNano() + 10)
	stateNow := <-state

	if stateNow == 0 {
		fmt.Println("Player2 joga a bola")
	} else {
		if isScore(stateNow) {
			SCOREPLAYER2++
			fmt.Printf("PLAYER 2 SCORED (%d)\n", SCOREPLAYER1)
		}
	}
	chanceOnPlay := rand.Intn(1000) + 1
	fmt.Println("Chance da jogada:", chanceOnPlay)
	state <- chanceOnPlay
	time.Sleep(1)
}

// Verifica se ocorreu pontuação
func isScore(chance int) bool {
	if chance < 100 {
		return true
	}
	return false
}

// Inicia a partida
func startGame(channel chan int) {
	go initializer(channel)
	for {
		if SCOREPLAYER1 > 4 && SCOREPLAYER1 > SCOREPLAYER2+2 {
			fmt.Printf("PLAYER1 WINS!!! P1 (%d) Vs (%d) P2", SCOREPLAYER1, SCOREPLAYER2)
			break
		}
		if SCOREPLAYER2 > 4 && SCOREPLAYER2 > SCOREPLAYER1+2 {
			fmt.Printf("PLAYER2 WINS!!! P1 (%d) Vs (%d) P2", SCOREPLAYER1, SCOREPLAYER2)
			break
		}
		go player1(channel)
		go player2(channel)
		time.Sleep(time.Duration(math.Pow10(9)))
	}
}

// Apresenta a tela inicial para o usuário
func systemPresentation() {
	line := "------------------------------"
	fmt.Println(line)
	fmt.Println("Simulador de partidas de Tenis")
	fmt.Println(line)
	fmt.Println("Escolha uma das opções:")
	fmt.Println("1 - Jogar")
	fmt.Println("2 - Configurar")
	fmt.Println(line)
	fmt.Println("0 - Finalizar Programa")
	fmt.Println(line)
}

// Define as configurações da partida. (PARTE EXTRA)
func setConfiguration() {
	fmt.Println("Configurando a partida!")
}

// Definição de variáveis
var SCOREPLAYER1 = 0
var SCOREPLAYER2 = 0
var GAME = 1
var SET = 1
var MATCH = 1

func main() {
	channel := make(chan int)
	var option int = 10
	var menuControl = true

	// Loop responsável por permancer no programa
	for menuControl {
		systemPresentation()
		fmt.Scanln(&option)
		switch option {
		case 1:
			startGame(channel)
			menuControl = false
		case 2:
			setConfiguration()
			startGame(channel)
			menuControl = false
		case 0:
			fmt.Println("Finalizando o simulador.")
			menuControl = false
		default:
			fmt.Println("Valor inválido. Digite uma opção válida!")
		}
	}
	time.Sleep(1e9)
}
