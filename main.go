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
	chanceOnPlay := rand.Intn(1000) + 1
	if stateNow == 0 {
		fmt.Println("PLAYER 1 SACA A BOLA")
	} else {
		if isScore(stateNow) {
			SCOREPLAYER1++
			fmt.Printf("PLAYER 1 SCORED (%d)\n", SCOREPLAYER1)
			chanceOnPlay = 0
		}
	}
	fmt.Println("Chance da jogada:", chanceOnPlay)
	state <- chanceOnPlay
}

// Goroutine do jogador2
func player2(state chan int) {
	fmt.Println("----------- PLAYER 2 -----------")
	rand.Seed(time.Now().UnixNano() + 10)
	stateNow := <-state
	chanceOnPlay := rand.Intn(1000) + 1
	if stateNow == 0 {
		fmt.Println("PLAYER 2 SACA A BOLA")
	} else {
		if isScore(stateNow) {
			SCOREPLAYER2++
			fmt.Printf("PLAYER 2 SCORED (%d)\n", SCOREPLAYER2)
			chanceOnPlay = 0
		}
	}
	fmt.Println("Chance da jogada:", chanceOnPlay)
	state <- chanceOnPlay
}

// Verifica se ocorreu pontuação
func isScore(chance int) bool {
	if chance < SUCESSCHANCE {
		return true
	}
	return false
}

// Inicia a partida
func startGame(channel chan int) {
	// Variável responsável por armazenar todas as pontuações do jogo
	scores := make([][][][]int, MATCH)

	for i := range scores {
		scores[i] = make([][][]int, SET)
		for j := range scores[i] {
			scores[i][j] = make([][]int, GAME)
			for k := range scores[i][j] {
				scores[i][j][k] = make([]int, 2)
			}
		}
	}

	go initializer(channel)
	for i := 1; i <= MATCH; i++ {
		fmt.Printf("MATCH %d STARTED!!!\n", i)
		SETPLAYER1 = 0
		SETPLAYER2 = 0
		for j := 1; j <= SET; j++ {
			fmt.Printf("SET %d STARTED!!!\n", j)
			if SETPLAYER1 >= (SET/2)+1 {
				fmt.Printf("PLAYER1 WINS MATCH %d!!! P1 (%d) Vs (%d) P2\n", i, SETPLAYER1, SETPLAYER2)
				break
			}
			if SETPLAYER2 >= (SET/2)+1 {
				fmt.Printf("PLAYER2 WINS MATCH %d!!! P1 (%d) Vs (%d) P2\n", i, SETPLAYER1, SETPLAYER2)
				break
			}
			GAMEPLAYER1 = 0
			GAMEPLAYER2 = 0
			for k := 1; k <= GAME; k++ {
				fmt.Printf("GAME %d STARTED!!!\n", k)
				if GAMEPLAYER1 >= (GAME/2)+1 {
					fmt.Printf("PLAYER1 WINS SET %d!!! P1 (%d) Vs (%d) P2\n", j, GAMEPLAYER1, GAMEPLAYER2)
					SETPLAYER1++
					break
				}
				if GAMEPLAYER2 >= (GAME/2)+1 {
					fmt.Printf("PLAYER2 WINS SET %d!!! P1 (%d) Vs (%d) P2\n", j, GAMEPLAYER1, GAMEPLAYER2)
					SETPLAYER2++
					break
				}
				SCOREPLAYER1 = 0
				SCOREPLAYER2 = 0
				for {
					if SCOREPLAYER1 >= SCORE && SCOREPLAYER1 > SCOREPLAYER2+2 {
						fmt.Printf("PLAYER1 WINS GAME %d!!! P1 (%d) Vs (%d) P2\n", k, SCOREPLAYER1, SCOREPLAYER2)
						scores[i-1][j-1][k-1][0] = SCOREPLAYER1
						scores[i-1][j-1][k-1][1] = SCOREPLAYER2
						GAMEPLAYER1++
						break
					}
					if SCOREPLAYER2 >= SCORE && SCOREPLAYER2 > SCOREPLAYER1+2 {
						fmt.Printf("PLAYER2 WINS GAME %d!!! P1 (%d) Vs (%d) P2\n", k, SCOREPLAYER1, SCOREPLAYER2)
						scores[i-1][j-1][k-1][0] = SCOREPLAYER1
						scores[i-1][j-1][k-1][1] = SCOREPLAYER2
						GAMEPLAYER2++
						break
					}
					go player1(channel)
					go player2(channel)
					time.Sleep(time.Duration(math.Pow10(SLEEPINGTIME)))
				}
				fmt.Printf("GAME %d ENDED!!!\n", k)
			}
			fmt.Printf("SET %d ENDED!!!\n", j)
		}
		fmt.Printf("MATCH %d ENDED!!!\n", i)
	}
	printScores(scores)
}

// Imprime o scoreboard final da partida
func printScores(scores [][][][]int) {
	fmt.Println("-----[SCORE BOARDING]-----")
	for i := range scores {
		fmt.Printf("----- MATCH %d -----\n", i+1)
		for j := range scores[i] {
			fmt.Printf("----- SET %d -----\n", j+1)
			for k := 1; k <= GAME; k++ {
				fmt.Printf("[%d] ", k)
			}
			fmt.Println("GAME Nº")
			for k := range scores[i][j] {
				if scores[i][j][k][0] == scores[i][j][k][1] {
					fmt.Printf("|-| ")
				} else {
					fmt.Printf("|%d| ", scores[i][j][k][0])
				}

			}
			fmt.Println("PLAYER 1")
			for k := range scores[i][j] {
				if scores[i][j][k][0] == scores[i][j][k][1] {
					fmt.Printf("|-| ")
				} else {
					fmt.Printf("|%d| ", scores[i][j][k][1])
				}
			}
			fmt.Println("PLAYER 2")
		}
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
	fmt.Println("Digite a quantidade de pontos para definir um GAME: ")
	fmt.Scanln(&SCORE)
	fmt.Println("Digite a quantidade de GAMES para definir um SET: ")
	fmt.Scanln(&GAME)
	fmt.Println("Digite a quantidade de SET's para definir um MATCH: ")
	fmt.Scanln(&SET)
	fmt.Println("Digite a quantidade de MATCH's da simulação: ")
	fmt.Scanln(&MATCH)
	fmt.Println("Digite a chance de sucesso para marcar um ponto (2-999): ")
	fmt.Scanln(&SUCESSCHANCE)
	fmt.Println("Digite a velocidade do log no terminal (0-10) (menor = mais rapido): ")
	fmt.Scanln(&SLEEPINGTIME)
}

// Definição de variáveis
var SCORE = 4
var SCOREPLAYER1 = 0
var SCOREPLAYER2 = 0
var GAME = 5
var GAMEPLAYER1 = 0
var GAMEPLAYER2 = 0
var SET = 3
var SETPLAYER1 = 0
var SETPLAYER2 = 0
var MATCH = 1
var SUCESSCHANCE = 200 // Valor entre 2 a 999 (define a chance de marcar ponto)
var SLEEPINGTIME = 8   // Valor entre 0 a 10 (define o timer do log no terminar)

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
