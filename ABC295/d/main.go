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
	S := nextLine()

	count := 0
	for i := 0; i < len(S); i++ {
		base := make([]int, 10)
		v, _ := strconv.Atoi(S[i : i+1])
		base[v]++

		for j := i + 1; j < len(S); j++ {
			t, _ := strconv.Atoi(S[j : j+1])
			base[t]++

			if check(base) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func check(base []int) bool {
	for i := 0; i < len(base); i++ {
		if base[i]%2 != 0 {
			return false
		}
	}
	return true
}
