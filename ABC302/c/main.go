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
	S := []string{}

	for i := 0; i < l[0]; i++ {
		S = append(S, nextLine())
	}

	sort.Strings(S)
	for nextPermutation(sort.StringSlice(S)) {
		res := true
		for i := 0; i < l[0]-1; i++ {
			a := strings.Split(S[i], "")
			b := strings.Split(S[i+1], "")

			diff := 0
			for k := 0; k < l[1]; k++ {
				if a[k] != b[k] {
					diff++
				}
			}
			if diff != 1 {
				res = false
			}
		}

		if res {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
