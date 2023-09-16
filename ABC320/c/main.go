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
	N := scanInt()
	S := []string{}
	for i := 0; i < 3; i++ {
		S = append(S, nextLine())
	}

	num := make([][][]int, 3)
	for i, v := range S {
		num[i] = make([][]int, 10)
		for pos, c := range v {
			t, _ := strconv.Atoi(string(c))
			num[i][t] = append(num[i][t], pos)
			num[i][t] = append(num[i][t], pos+N)
			num[i][t] = append(num[i][t], pos+(N*2))
		}
	}

	min := 10000000
	for i := 0; i < 10; i++ {
		res := slot(num[0][i], num[1][i], num[2][i])
		if res > 0 && min > res {
			min = res
		}
	}
	if min == 10000000 {
		min = -1
	}
	fmt.Println(min)
}

func slot(s1 []int, s2 []int, s3 []int) int {
	sort.Ints(s1)
	sort.Ints(s2)
	sort.Ints(s3)
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			for k := 0; k < len(s3); k++ {
				if s1[i] != s2[j] && s1[i] != s3[k] && s2[j] != s3[k] {
					return max(s1[i], s2[j], s3[k])
				}
			}
		}
	}
	return -1
}

func max(s1, s2, s3 int) int {
	res := s1
	if res < s2 {
		res = s2
	}

	if res < s3 {
		res = s3
	}
	return res
}
