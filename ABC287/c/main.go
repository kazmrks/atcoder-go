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

	var N, M int
	fmt.Scan(&N, &M)

	if N != M+1 {
		fmt.Println("No")
		return
	}

	path := map[int][]int{}
	for i := 0; i < M; i++ {
		node := lineToInts(2)
		u := node[0]
		v := node[1]

		p, exists := path[u]
		if !exists {
			path[u] = []int{v, 0}
		} else {
			if p[1] != 0 || p[0] == v {
				fmt.Println("No")
				return
			}
			path[u] = []int{p[0], v}
		}

		p, exists = path[v]
		if !exists {
			path[v] = []int{u, 0}
		} else {
			if p[1] != 0 || p[0] == u {
				fmt.Println("No")
				return
			}
			path[v] = []int{p[0], u}
		}
	}

	edge := 0
	for i := 1; i <= N; i++ {
		p, exists := path[i]

		if !exists {
			fmt.Println("No")
			return
		}

		if p[1] == 0 {
			edge++
		}
	}

	if edge == 2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
