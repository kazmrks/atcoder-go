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

	var N, M int64
	fmt.Scan(&N, &M)

	if N*N < M {
		fmt.Println("-1")
		return
	}

	check := N * N
	for b := int64(1); b <= N; b++ {
		a := math.Ceil(float64(M) / float64(b))

		if int64(a) > N {
			continue
		}

		ab := int64(a) * b
		if math.Floor(a) == a && ab >= M {
			if check >= ab {
				check = ab
			}
		}
	}

	if check <= N*N {
		fmt.Println(check)
		return
	}

	fmt.Println("-1")
}
