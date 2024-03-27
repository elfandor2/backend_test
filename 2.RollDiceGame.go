package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Dices []int
	Score int
}

func rollDiceGame(n, m int) {
	var players []Player
	var playerScore []int
	var count int
	var winner int
	var playerNumber int

	for i := 0; i < n; i++ {
		player := Player{
			Dices: make([]int, 0),
			Score: 0,
		}

		for j := 0; j < m; j++ {
			diceValue := rand.Intn(6) + 1
			player.Dices = append(player.Dices, diceValue)
		}
		players = append(players, player)
	}

	fmt.Println("Lempar dadu")
	for i, player := range players {
		fmt.Printf("Player #%d (%d): %v\n", i+1, player.Score, player.Dices)
	}

	for {
		players = evaluateAndRollDice(players)
		activePlayers := 0
		for _, player := range players {
			if len(player.Dices) > 0 {
				activePlayers++
			}
		}
		if activePlayers <= 1 {
			break
		}
	}

	fmt.Println()
	fmt.Println("Game berkahir, hanya tersisa satu pemain:")
	for i, player := range players {
		fmt.Printf("Player #%d (%d): %v\n", i+1, player.Score, player.Dices)
	}

	for i := 0; i < len(players); i++ {
		if winner < players[i].Score {
			winner = players[i].Score
			count = 1
			playerNumber = i + 1
		} else if winner == players[i].Score {
			count++
		}
		playerScore = append(playerScore, players[i].Score)

	}
	if count == 1 {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", playerNumber)
	} else {
		fmt.Printf("Game berakhir seri.\n")
	}
	fmt.Println(playerScore)

}

func rollDice(players []Player) []Player {
	for i := 0; i < len(players); i++ {
		var newDices []int

		for j := 0; j < len(players[i].Dices); j++ {
			diceValue := rand.Intn(6) + 1
			newDices = append(newDices, diceValue)
		}

		players[i].Dices = newDices
	}

	fmt.Println("Lempar dadu")
	for i, player := range players {
		fmt.Printf("Player #%d (%d): %v\n", i+1, player.Score, player.Dices)
	}

	return players
}

func evaluateRollDice(players []Player) []Player {
	for i := 0; i < len(players); i++ {
		var newDices []int

		for j := 0; j < len(players[i].Dices); j++ {
			diceValue := players[i].Dices[j]

			if diceValue == 1 && i < len(players)-1 {
				players[i+1].Dices = append(players[i+1].Dices, 7)
			}

			if diceValue != 6 && diceValue != 1 {
				newDices = append(newDices, diceValue)
			} else if diceValue == 6 {
				players[i].Score++
			} else if diceValue == 1 {
				if i+1 == len(players) {
					players[0].Dices = append(players[0].Dices, 1)
				}
			}
		}

		players[i].Dices = newDices
	}

	fmt.Println("Setelah Evaluasi:")
	for i, player := range players {
		fmt.Printf("Player #%d (%d): %v\n", i+1, player.Score, player.Dices)
	}
	fmt.Println()

	return players
}

func evaluateAndRollDice(players []Player) []Player {
	evaluatedPlayers := evaluateRollDice(players)
	rolledPlayers := rollDice(evaluatedPlayers)
	return rolledPlayers
}

func main() {
	rollDiceGame(3, 4)
}
