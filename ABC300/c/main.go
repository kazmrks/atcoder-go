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

func main() {
	line := lineToInts()
	H := line[0]
	W := line[1]

	C := [][]rune{}

	for i := 0; i < H; i++ {
		C = append(C, []rune(nextLine()))
	}

	S := make([]int, min(H, W))

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if C[i][j] == '#' {
				if test(C, i, j, H, W, 1) {
					n := 1
					for test(C, i, j, H, W, n+1) {
						n++
					}
					S[n-1]++
				}
			}
		}
	}
	printInts(S)
}

func min(H, W int) int {
	if H < W {
		return H
	} else {
		return W
	}
}

func test(C [][]rune, i, j, H, W, d int) bool {
	if i-d < 0 || j-d < 0 || i+d >= H || j+d >= W {
		return false
	}

	if C[i-d][j-d] == '.' || C[i-d][j+d] == '.' || C[i+d][j-d] == '.' || C[i+d][j+d] == '.' {
		return false
	}
	return true
}
