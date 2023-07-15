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

	S := nextLine()
	T := nextLine()

	s1 := [27]int{}
	s2 := [27]int{}

	for i := 0; i < len(S); i++ {
		if S[i] == '@' {
			s1[0]++
		} else {
			pos := S[i] - 'a' + 1
			s1[pos]++
		}

		if T[i] == '@' {
			s2[0]++
		} else {
			pos := T[i] - 'a' + 1
			s2[pos]++
		}
	}

	d1 := 0
	d2 := 0
	for i := 1; i < 26; i++ {
		if s1[i] > s2[i] {
			if notAtcoder(i) {
				fmt.Println("No")
				return
			}
			d1 = d1 + s1[i] - s2[i]
		}
		if s1[i] < s2[i] {
			if notAtcoder(i) {
				fmt.Println("No")
				return
			}
			d2 = d2 + s2[i] - s1[i]
		}
	}

	if s1[0] >= d2 {
		fmt.Println("Yes")
		return
	} else {
		fmt.Println("No")
	}
}

func notAtcoder(v int) bool {
	if v == 1 || v == 3 || v == 4 || v == 5 || v == 15 || v == 18 || v == 20 {
		return false
	}
	return true
}
