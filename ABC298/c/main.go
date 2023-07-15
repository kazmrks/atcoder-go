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
	Q := scanInt()

	two := map[int][]int{}
	three := map[int][]int{}
	for i := 1; i <= Q; i++ {
		q := lineToInts()

		switch q[0] {
		case 1:
			two[q[2]] = append(two[q[2]], q[1])
			if !contains(three[q[1]], q[2]) {
				three[q[1]] = append(three[q[1]], q[2])
			}
		case 2:
			sort.Ints(two[q[1]])
			print(two[q[1]])
		case 3:
			sort.Ints(three[q[1]])
			print(three[q[1]])
		}
	}
}
