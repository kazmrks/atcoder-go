package main

import (
	"bufio"
	"fmt"
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

func main() {
	S := nextLine()
	b := []byte(S)

	for i := 1; i <= len(S)/2; i++ {
		buf := b[(2*i)-1-1]
		b[(2*i)-1-1] = b[2*i-1]
		b[2*i-1] = buf
	}
	fmt.Println(string(b))
}
