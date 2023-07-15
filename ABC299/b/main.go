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

func main() {
	sc.Buffer([]byte{}, math.MaxInt64)

	L := lineToInts()
	N, T := L[0], L[1]

	C := lineToInts()
	R := lineToInts()

	p1, pt := 0, 0
	win1, wint := 0, 0
	flag := false

	for i := 0; i < N; i++ {
		if C[i] == T && R[i] > pt {
			pt = R[i]
			wint = i + 1
			flag = true
		}

		if C[i] == C[0] && R[i] > p1 {
			p1 = R[i]
			win1 = i + 1
		}
	}

	if flag {
		fmt.Println(wint)
	} else {
		fmt.Println(win1)
	}
}
