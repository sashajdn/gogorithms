package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Colour int

const (
	ColourYellow Colour = iota + 1
	ColourRed
)

func (c Colour) String() string {
	switch c {
	case ColourRed:
		return "R"
	case ColourYellow:
		return "Y"
	default:
		return ""
	}
}

type ConnectFourEntryType int

const (
	ConnectFourEntryTypeEmpty ConnectFourEntryType = iota
	ConnectFourEntryTypeRed
	ConnectFourEntryTypeYellow
)

type ConnectFourEntry struct {
	column int
	player *Player
}

func (c *ConnectFourEntry) Validate(board ConnectFourBoard) error {
	if c.column < 0 || c.column >= board.Width() {
		return fmt.Errorf("column invalid: %d outside board bounds", c.column+1)
	}

	return nil
}

func (c ConnectFourEntryType) String() string {
	switch c {
	case ConnectFourEntryTypeEmpty:
		return " "
	case ConnectFourEntryTypeYellow:
		return fmt.Sprintf("%s%s\033[0m", "\033[33m", "O")
	case ConnectFourEntryTypeRed:
		return fmt.Sprintf("%s%s\033[0m", "\033[31m", "O")
	default:
		return " "
	}
}

type Player struct {
	Name   string
	Colour Colour
}

func (p *Player) AskTurn() *ConnectFourEntry {
	var column int
	fmt.Printf("\n%s (%s) please pick column: ", p.Name, p.Colour)
	fmt.Scan(&column)

	return &ConnectFourEntry{
		column: column - 1,
		player: p,
	}
}

func (p *Player) AskName() {
	var name string

	fmt.Printf("\n%s player please state your name: ", p.Colour)
	fmt.Scan(&name)

	p.Name = name
}

func NewConnectFourGame(playerOne, playerTwo *Player, board ConnectFourBoard) *ConnectFour {
	var firstTurn = playerOne
	if rand.Float64() > 0.5 {
		firstTurn = playerTwo
	}

	availableRows := make([]int, board.Width())
	for i := 0; i < board.Height(); i++ {
		availableRows[i] = board.Height() - 1
	}

	return &ConnectFour{
		board:         board,
		playerOne:     playerOne,
		playerTwo:     playerTwo,
		availableRows: availableRows,
		currentTurn:   firstTurn,
	}
}

type ConnectFour struct {
	board                ConnectFourBoard
	availableRows        []int
	playerOne, playerTwo *Player
	currentTurn          *Player
	lastEntry            *ConnectFourEntry
	numberOfTurns        int
}

func (c *ConnectFour) NextTurn() *Player {
	if c.numberOfTurns > c.board.Width()*c.board.Height() {
		return nil
	}

	if c.currentTurn == c.playerOne {
		c.currentTurn = c.playerTwo
		return c.playerTwo
	}

	c.currentTurn = c.playerOne
	return c.playerOne
}

func (c *ConnectFour) PlayTurn(entry *ConnectFourEntry) {
	column := entry.column
	row := c.availableRows[column]

	var boardEntry = ConnectFourEntryTypeRed
	if entry.player.Colour == ColourYellow {
		boardEntry = ConnectFourEntryTypeYellow
	}

	c.board[row][column] = boardEntry
	c.availableRows[column]--
	c.lastEntry = entry
}

func (c *ConnectFour) hasWon() bool {
	winningString := strings.Repeat(c.lastEntry.player.Colour.String(), 4)

	if c.winningNegativeDiagonal(winningString) || c.winningPositiveDiagonal(winningString) || c.winningVertical(winningString) || c.winningHorizontal(winningString) {
		return true
	}

	return false
}

func (c *ConnectFour) winningVertical(winningString string) bool {
	lastColumn := c.lastEntry.column
	lastRow := c.availableRows[lastColumn]

	var s string
	for row := 0; row < c.board.Height() && row >= lastRow+3; row++ {
		s += c.board[row][lastColumn].String()
	}

	return s == winningString
}

func (c *ConnectFour) winningHorizontal(winningString string) bool {
	lastColumn := c.lastEntry.column

	l, r := lastColumn-3, lastColumn
	for l < 0 {
		l++
	}

	for r-l < 4 {
		r++
	}

	for r < lastColumn+3 {
		slidingWindow := c.board[c.availableRows[lastColumn]][l:r]

		var s string
		for _, v := range slidingWindow {
			s += v.String()
		}

		if s == winningString {
			return true
		}

		r++
	}

	return false
}

func (c *ConnectFour) winningPositiveDiagonal(winningString string) bool {
	lastColumn := c.lastEntry.column
	lastRow := c.availableRows[lastColumn]

	var s string

	col, row := lastColumn, lastRow
	for col < c.board.Width() && row < c.board.Height() {
		s += c.board[row][col].String()
		col++
		row++
	}

	return s == winningString
}

func (c *ConnectFour) winningNegativeDiagonal(winningString string) bool {
	lastColumn := c.lastEntry.column
	lastRow := c.availableRows[lastColumn]

	var s string

	col, row := lastColumn, lastRow
	for col >= 0 && row < c.board.Height() {
		s += c.board[row][col].String()
		row++
		col--
	}

	return s == winningString
}

func NewConnectFourBoard(width, height int) ConnectFourBoard {
	var board = make(ConnectFourBoard, 0, height)
	for i := 0; i < height; i++ {
		board = append(board, make([]ConnectFourEntryType, width))
	}

	return board
}

type ConnectFourBoard [][]ConnectFourEntryType

func (c ConnectFourBoard) Width() int {
	if len(c) > 0 {
		return len(c[0])
	}

	return 0
}

func (c ConnectFourBoard) Height() int {
	return len(c)
}

func (c ConnectFourBoard) Display() {
	var columnHeader string
	for i := 0; i < c.Width(); i++ {
		columnHeader += fmt.Sprintf("| %d |", i+1)
	}
	fmt.Println(columnHeader)
	fmt.Println()

	for j := 0; j < c.Height(); j++ {
		var s string
		for i := 0; i < c.Width(); i++ {
			s += fmt.Sprintf("| %s |", c[j][i].String())
		}

		fmt.Println(s)
	}
}
