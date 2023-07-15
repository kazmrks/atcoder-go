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

	l := lineToInts()
	H := l[0]
	W := l[1]

	A := [][]rune{}
	B := [][]rune{}

	for i := 0; i < H; i++ {
		line := nextLine()
		r := []rune(line)
		A = append(A, r)
	}

	for i := 0; i < H; i++ {
		line := nextLine()
		r := []rune(line)
		B = append(B, r)
	}

	for s := 0; s < H; s++ {
		for t := 0; t < W; t++ {

			match := true
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					if A[(i+s)%H][(j+t)%W] != B[i][j] {
						match = false
						break
					}
				}
			}
			if match {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}
