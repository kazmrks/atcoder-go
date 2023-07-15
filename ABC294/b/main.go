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

	var H, W int
	fmt.Scan(&H, &W)

	A := [][]int{}
	for i := 0; i < H; i++ {
		A = append(A, lineToInts(W))
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if A[i][j] == 0 {
				fmt.Print(".")

			} else {
				r := 65
				fmt.Printf("%c", r-1+A[i][j])
			}
		}
		fmt.Printf("\n")
	}
}
