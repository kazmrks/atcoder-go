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

	var N int
	fmt.Scan(&N)

	S := nextLine()

	bar := []int{}

	for i, v := range S {
		if i == 0 {
			if v == '-' {
				bar = append(bar, 0)
			} else {
				bar = append(bar, -1)
			}
			continue
		}

		if v == '-' {
			bar = append(bar, i)
		}
	}

	if len(bar) == 1 && bar[0] == -1 {
		fmt.Println("-1")
		return
	}

	if S[N-1] == 'o' {
		bar = append(bar, N)
	}

	res := 0
	x := 0
	for i, v := range bar {
		if i == 0 {
			x = v
			continue
		}

		tmp := v - x - 1
		if tmp > res {
			res = tmp
		}
		x = v
	}

	if res == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(res)
	}
}
