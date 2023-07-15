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
	called := make([]bool, N)

	for i := 0; i < N; i++ {
		if !called[i] {
			called[A[i]-1] = true
		}
	}

	X := []string{}
	for i := 0; i < N; i++ {
		if !called[i] {
			X = append(X, strconv.Itoa(i+1))
		}
	}

	fmt.Println(len(X))
	fmt.Println(strings.Join(X, " "))
}
