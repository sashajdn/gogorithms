package main

import "fmt"

func main() {
	playerOne := &Player{
		Colour: ColourRed,
	}

	playerTwo := &Player{
		Colour: ColourYellow,
	}

	width, height := readBoardDimensions()

	playerOne.AskName()
	playerTwo.AskName()

	board := NewConnectFourBoard(width, height)

	game := NewConnectFourGame(playerOne, playerTwo, board)
	var winner *Player

	fmt.Printf("\n\033[36m=================== CONNECT 4 **********************\033[0m\n\n")
	fmt.Printf("            >>>>>> %s vs %s <<<<<<<\n\n", playerOne.Name, playerTwo.Name)
	for {
		board.Display()

		if player := game.NextTurn(); player != nil {
			var entry *ConnectFourEntry
			for {
				entry = player.AskTurn()
				if err := entry.Validate(game.board); err != nil {
					fmt.Printf("\nInvalid entry: %d, please try again...\n", entry.column+1)
					continue
				}
				break
			}

			game.PlayTurn(entry)
			if game.hasWon() {
				winner = player
				break
			}

			continue
		}

		break
	}

	if winner == nil {
		fmt.Println("No winner for this game...")
		return
	}

	fmt.Printf("\nCongrats: %s has won!!\n", winner.Name)
}

func readBoardDimensions() (int, int) {
	var w, h int

	for {
		fmt.Printf("\nPlease choose the height of the board: ")
		if _, err := fmt.Scan(&h); err != nil {
			fmt.Println("Please choose valid height: must be an integer...")
			continue
		}

		break
	}

	for {
		fmt.Printf("\nPlease choose the width of the board: ")
		if _, err := fmt.Scan(&w); err != nil {
			fmt.Println("Please choose valid width: must be an integer...")
			continue
		}

		break
	}

	return w, h
}
