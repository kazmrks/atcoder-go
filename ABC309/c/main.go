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

type pair struct {
	a int
	b int
}

func main() {
	line := lineToInts()
	N, K := line[0], line[1]

	med := []pair{}
	count := 0

	for i := 0; i < N; i++ {
		l := lineToInts()
		a, b := l[0], l[1]
		med = append(med, pair{a, b})
		count += b
	}

	sort.SliceStable(med, func(i, j int) bool {
		return med[i].a < med[j].a
	})

	if count <= K {
		fmt.Println(1)
		return
	}

	for i := 0; i < len(med); i++ {
		count -= med[i].b
		if count <= K {
			fmt.Println(med[i].a + 1)
			return
		}
	}

	fmt.Println(med[len(med)-1].a + 1)
}
