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
	l := lineToInts()
	N := l[0]

	items := []item{}
	for i := 0; i < N; i++ {
		line := lineToInts()
		items = append(items, item{P: line[0], C: line[1], F: line[2:]})
	}

	for i := 0; i < N; i++ {
		item_i := items[i]
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			item_j := items[j]

			if item_i.P < item_j.P || item_i.C > item_j.C {
				continue
			}

			ok := true
			for _, v := range item_i.F {
				if !contains(item_j.F, v) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}

			if item_i.P > item_j.P || item_i.C < item_j.C {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

type item struct {
	P int
	C int
	F []int
}
