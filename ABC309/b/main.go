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
	N := scanInt()

	A := [][]int{}
	B := [][]int{}
	for i := 0; i < N; i++ {
		line := nextLine()
		a := make([]int, N)
		b := make([]int, N)

		for j := 0; j < N; j++ {
			v, _ := strconv.Atoi(string(line[j]))
			a[j] = v
			b[j] = v
		}
		A = append(A, a)
		B = append(B, b)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == 0 {
				if j == 0 {
					B[i][j] = A[1][0]
				} else {
					B[i][j] = A[i][j-1]
				}
				continue
			}

			if i == N-1 {
				if j == N-1 {
					B[i][j] = A[i-1][j]
				} else {
					B[i][j] = A[i][j+1]
				}
				continue
			}

			if j == 0 {
				B[i][j] = A[i+1][j]
				continue
			}

			if j == N-1 {
				B[i][j] = A[i-1][j]
				continue
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(B[i][j])
		}
		fmt.Print("\n")
	}
}
