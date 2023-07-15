package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	i, err := strconv.Atoi(nextLine())
	if err != nil {
		panic(err)
	}
	return i
}

func lineToInts(n int) []int {
	slice := make([]int, n)
	text := strings.Split(nextLine(), " ")

	for i := 0; i < n; i++ {
		v, err := strconv.Atoi(text[i])
		if err != nil {
			panic(err)
		}
		slice[i] = v
	}
	return slice
}

func contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	sc.Buffer([]byte{}, math.MaxInt64)

	var R, C int
	fmt.Scan(&R, &C)

	board := make([][]string, R)
	bomb := make([][]string, R)
	for i := 0; i < R; i++ {
		board[i] = make([]string, C)
		bomb[i] = make([]string, C)
	}

	for i := 0; i < R; i++ {
		B := nextLine()
		for j := 0; j < C; j++ {
			board[i][j] = B[j : j+1]
			bomb[i][j] = B[j : j+1]
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			switch bomb[i][j] {
			case ".", "#":
			default:
				power, _ := strconv.Atoi(bomb[i][j])
				board = bombDown(board, i, j, R, C, power)
				board = bombUp(board, i, j, R, C, power)
			}
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Print("\n")
	}
}

func bombDown(board [][]string, i int, j int, R int, C int, power int) [][]string {
	for x := j; x >= j-power; x-- {
		if x >= 0 {
			board[i][x] = "."
		}
	}

	for x := j; x <= j+power; x++ {
		if x < C {
			board[i][x] = "."
		}
	}

	if power-1 < 0 || i-1 < 0 {
		return board
	} else {
		return bombDown(board, i-1, j, R, C, power-1)
	}
}

func bombUp(board [][]string, i int, j int, R int, C int, power int) [][]string {
	for x := j; x >= j-power; x-- {
		if x >= 0 {
			board[i][x] = "."
		}
	}

	for x := j; x <= j+power; x++ {
		if x < C {
			board[i][x] = "."
		}
	}

	if power-1 < 0 || i+1 == R {
		return board
	} else {
		return bombUp(board, i+1, j, R, C, power-1)
	}
}
