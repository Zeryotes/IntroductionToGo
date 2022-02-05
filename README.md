# IntroductionToGo
Introdução a linguagem de programação Golang (Go) e os mecanismos de programação concorrente presentes na linguagem.

# Problemática
Desenvolver um programa em Golang que represente um jogo de tênis com dois jogadores (problema 1 do main.pdf).

# Orientações
## Prints nas goroutines
Nas funções executadas pelas goroutines (player1 e player2) existem alguns "fmt.Println()" comentados. Utilizei para acompanhar o funcionamento interno de cada goroutine (ação de cada player). De acordo com o que será observado será interessante comentar/descomentar esses prints. 
## Tempo de espera
Recomendo definir o SLEEPINGTIME de 8 ou 9 para acompanhamento das ações de cada jogador.
Recomendo definri o SLEEPINGTIME de 5 a 7 para acompanhamento dos conjuntos MATCH-SET-GAME.
## Definição do sucesso das jogadas
Essa definição está sendo realizada por meio da variável SUCESSCHANCE. Está sendo gerado um número aleatório entre 1 e 1000. O SUCESSCHANCE é o range dentre 1 a SUCESSCHANCE que define se a jogada foi bem sucessidade ou não. 
