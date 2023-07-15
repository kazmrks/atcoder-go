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

	N := scanInt()
	A := lineToInts()
	B := []int{A[0]}
	for i := 1; i < N; i++ {
		ai := B[len(B)-1]
		if math.Abs(float64(ai-A[i])) == 1 {
			B = append(B, A[i])
			continue
		}

		if ai < A[i] {
			diff := A[i] - ai
			for j := 1; j < diff; j++ {
				B = append(B, ai+j)
			}
		} else if ai > A[i] {
			diff := ai - A[i]
			for j := 1; j < diff; j++ {
				B = append(B, ai-j)
			}
		}
		B = append(B, A[i])
	}
	printInts(B)
}
