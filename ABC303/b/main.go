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

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

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

	for _, t := range text {
		v, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		slice = append(slice, v)
	}
	return slice
}

func printInts(s []int) {
	p := []string{}
	for _, v := range s {
		p = append(p, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(p, " "))
}

func printStrings(s []string) {
	fmt.Println(strings.Join(s, " "))
}

func contains(elements []int, v int) bool {
	for _, e := range elements {
		if e == v {
			return true
		}
	}
	return false
}

// 順列全探索
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	line := lineToInts()
	N := line[0]
	M := line[1]

	a := [][]int{}
	for i := 0; i < M; i++ {
		a = append(a, lineToInts())
	}

	p := [][]int{}
	for i := 0; i < N; i++ {
		p = append(p, make([]int, N))
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N-1; j++ {
			s := a[i][j] - 1
			t := a[i][j+1] - 1
			if s > t {
				v := s
				s = t
				t = v
			}
			p[s][t]++
		}
	}

	res := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if p[i][j] == 0 {
				res++
			}
		}
	}
	fmt.Println(res)
}
