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

	var i, j int
	check := false
	for i = 0; i < 8; i++ {
		S := nextLine()
		ary := strings.Split(S, "")
		for j = 0; j < 8; j++ {
			if ary[j] == "*" {
				check = true
				break
			}
		}
		if check {
			break
		}
	}

	res := ""
	switch j {
	case 0:
		res = "a"
	case 1:
		res = "b"
	case 2:
		res = "c"
	case 3:
		res = "d"
	case 4:
		res = "e"
	case 5:
		res = "f"
	case 6:
		res = "g"
	case 7:
		res = "h"
	}

	res = res + strconv.Itoa(8-i)

	fmt.Println(res)
}
