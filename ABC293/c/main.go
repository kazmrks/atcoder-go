package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
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

func search(H int, W int, i int, j int, p [][]int, taka []int) (result int) {
	if contains(taka, p[i-1][j-1]) {
		return 0
	}

	if i == H && j == W {
		return 1
	}

	taka = append(taka, p[i-1][j-1])

	if j < W {
		result = result + search(H, W, i, j+1, p, taka)
	}

	if i < H {
		result = result + search(H, W, i+1, j, p, taka)
	}
	return
}

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	A := [][]int{}
	for i := 1; i <= H; i++ {
		A = append(A, lineToInts(W))
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

		}
	}
	fmt.Println(search(H, W, 1, 1, A, []int{}))
}
