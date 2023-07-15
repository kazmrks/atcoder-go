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

func main() {
	sc.Buffer([]byte{}, math.MaxInt64)

	N := scanInt()
	A := lineToInts(N)
	M := scanInt()
	B := lineToInts(M)
	X := scanInt()

	steps := make([]bool, X+1)
	steps[0] = true

	mochi := make(map[int]bool, M)
	for i := 0; i < M; i++ {
		mochi[B[i]] = true
	}

	for i := 0; i < X; i++ {
		if !steps[i] {
			continue
		}

		for j := 0; j < N; j++ {
			next := i + A[j]
			_, m := mochi[next]
			if next <= X && !m {
				steps[next] = true
			}
		}
	}

	if steps[X] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
