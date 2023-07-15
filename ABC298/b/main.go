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

	N := scanInt()

	A := [][]int{}
	for i := 0; i < N; i++ {
		A = append(A, lineToInts(N))
	}

	B := [][]int{}
	for i := 0; i < N; i++ {
		B = append(B, lineToInts(N))
	}

	ok := true
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i][j] == 1 && B[i][j] == 0 {
				ok = false
				break
			}
		}
	}
	if ok {
		fmt.Println("Yes")
		return
	}

	for count := 1; count < 4; count++ {
		ok = true
		T := [][]int{}
		for i := 0; i < N; i++ {
			a := make([]int, len(A[i]))
			copy(a, A[i])
			T = append(T, a)
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				A[i][j] = T[N-1-j][i]
			}
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if A[i][j] == 1 && B[i][j] == 0 {
					ok = false
					break
				}
			}
		}
		if ok {
			fmt.Println("Yes")
			return
		}

	}
	fmt.Println("No")
}
