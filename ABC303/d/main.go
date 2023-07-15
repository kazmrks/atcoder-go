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

// 順列全探索
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	line := lineToInts()
	X := line[0]
	Y := line[1]
	Z := line[2]
	zx := Z + X
	zy := Z + Y

	S := nextLine()
	e := make([]int, len(S))

	caps := false
	dup := 1
	for i := 0; i < len(S); i++ {
		if i == 0 {
			if S[i] == 'a' {
				e[0] = X
				if X > zy {
					e[0] = zy
					caps = true
				}
			}

			if S[i] == 'A' {
				e[0] = Y
				if Y > zx {
					e[0] = zx
					caps = true
				}
			}
			continue
		}

		if S[i-1] == S[i] {
			dup++
		} else {
			dup = 1
		}

		if S[i] == 'a' {
			if !caps {
				e[i] = e[i-1] + X

				if dup >= 2 && e[i-1]+Y < e[i] {
					e[i] = e[i-1] + Y
				}

				if X > zy {
					e[i] = e[i-1] + zy
					caps = true
				}
			} else {
				e[i] = e[i-1] + Y

				if dup >= 2 && e[i-1]+X < e[i] {
					e[i] = e[i-1] + X
				}

				if Y > zx {
					e[i] = e[i-1] + zx
					caps = false
				}
			}
		} else if S[i] == 'A' {
			if !caps {
				e[i] = e[i-1] + Y

				if dup >= 2 && e[i-1]+X < e[i] {
					e[i] = e[i-1] + X
				}

				if Y > zx {
					e[i] = e[i-1] + zx
					caps = true
				}
			} else {
				e[i] = e[i-1] + X

				if dup >= 2 && e[i-1]+Y < e[i] {
					e[i] = e[i-1] + Y
				}

				if X > zy {
					e[i] = e[i-1] + zy
					caps = false
				}
			}
		}
	}
}
