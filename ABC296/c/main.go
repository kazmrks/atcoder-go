package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	var N, X int
	fmt.Scan(&N, &X)

	if X == 0 {
		fmt.Println("Yes")
		return
	}

	A := lineToInts(N)
	sort.Ints(A)

	m := map[int]bool{}
	t := []int{}

	for _, a := range A {
		if !m[a] {
			m[a] = true
			t = append(t, a)
		}
	}

	for _, a := range t {
		res := X + a
		if m[res] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
