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

	var N, M int
	fmt.Scan(&N, &M)

	A := lineToInts(N)
	B := lineToInts(M)

	C := append(A, B...)
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	i := 0
	K1 := []string{}
	for k, v := range C {
		if i == N {
			break
		}

		if v == A[i] {
			K1 = append(K1, strconv.Itoa(k+1))
			i++
		}
	}

	i = 0
	K2 := []string{}
	for k, v := range C {
		if i == M {
			break
		}

		if v == B[i] {
			K2 = append(K2, strconv.Itoa(k+1))
			i++
		}
	}

	fmt.Println(strings.Join(K1, " "))
	fmt.Println(strings.Join(K2, " "))
}
