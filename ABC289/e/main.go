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

func scanInt() int {
	i, err := strconv.Atoi(nextLine())
	if err != nil {
		panic(err)
	}
	return i
}

func lineToInts() []int {
	slice := []int{}
	text := strings.Split(nextLine(), " ")

	for _, v := range text {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		slice = append(slice, i)
	}
	return slice
}

type testCase struct {
	N int
	M int
	C []int
	P []pair
}

type pair struct {
	u int
	v int
}

func main() {
	T := scanInt()

	for i := 1; i <= T; i++ {
		var test testCase

		v := lineToInts()
		test.N = v[0]
		test.M = v[1]
		test.C = lineToInts()
		test.P = []pair{}

		for j := 1; j <= test.M; j++ {
			p := lineToInts()
			test.P = append(test.P, pair{p[0], p[1]})
		}

		calcTest(test)
	}
}

func calcTest(test testCase) {
	fmt.Println(test)

	// 到達可能なルート情報を保持
	route := map[int][]int{}
	for _, p := range test.P {
		if !contains(route[p.u], p.v) {
			route[p.u] = append(route[p.u], p.v)
		}
		if !contains(route[p.v], p.u) {
			route[p.v] = append(route[p.v], p.u)
		}
	}

}

func contains(route []int, to int) bool {
	for _, v := range route {
		if v == to {
			return true
		}
	}
	return false
}
