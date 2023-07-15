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

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var C []int
	var S [][]int
	for i := 0; i < M; i++ {
		c, _ := strconv.Atoi(nextLine())
		C = append(C, c)

		text := strings.Split(nextLine(), " ")
		var s []int
		for _, v := range text {
			a, _ := strconv.Atoi(v)
			s = append(s, a)
		}
		S = append(S, s)
	}

	result := 0
	for i := 0; i < M; i++ {
		result = result + counter(N, S[i:i+1], S[i+1:], 0)
	}
	fmt.Println(result)

}

func counter(N int, s1 [][]int, s2 [][]int, count int) int {
	mixed := []int{}
	for _, s := range s1 {
		for _, v := range s {
			if !contains(mixed, v) {
				mixed = append(mixed, v)
			}
		}
	}

	if len(mixed) == N {
		count++
	}

	for i := 0; i < len(s2); i++ {
		t := [][]int{}
		t = append(t, s1...)
		t = append(t, s2[i])
		count = count + counter(N, t, s2[i+1:], 0)
	}

	return count
}

func contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}
