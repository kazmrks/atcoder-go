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

	X := []int{}
	que := []int{}
	n := 1
	for i := 0; i < M; i++ {
		q := strings.Split(nextLine(), " ")

		if q[0] == "1" {
			que = append(que, n)
			n++
		} else if q[0] == "2" {
			for j := 0; j < len(que); j++ {
				if q[1] == strconv.Itoa(que[j]) {
					que = append(que[:j], que[j+1:]...)
				}
			}
		} else if q[0] == "3" {
			if len(que) > 0 {
				X = append(X, que[0])
			}
		}
	}

	for _, v := range X {
		fmt.Println(v)
	}
}
