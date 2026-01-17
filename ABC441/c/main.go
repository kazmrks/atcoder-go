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

func lineToInts64() []int64 {
	slice := []int64{}
	text := strings.Split(nextLine(), " ")

	for _, t := range text {
		v, err := strconv.ParseInt(t, 10, 64)
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
	line := lineToInts64()
	N, K, X := line[0], line[1], line[2]

	A := lineToInts64()

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	total := int64(0)
	min := int64(0)
	solo := false
	for i, jp := int64(0), int64(0); i < K; i++ {
		total += A[i]

		if A[i] >= X {
			solo = true
		}
		if jp < X {
			jp += A[i]
			min++
		}
	}

	if total < X && !solo {
		fmt.Println(-1)
		return
	}

	if min == K {
		fmt.Println(N)
		return
	}

	fmt.Println(min + N - K)

}
