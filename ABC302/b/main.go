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
	l := lineToInts()
	S := [][]rune{}

	for i := 0; i < l[0]; i++ {
		S = append(S, []rune(nextLine()))
	}

	for i := 0; i < l[0]; i++ {
		for j := 0; j < l[1]; j++ {
			if S[i][j] == 's' {
				ok, res := isSnuke(S, i, j, l[0], l[1])
				if ok {
					for _, v := range res {
						printInts(v)
					}
					return
				}
			}
		}
	}
}

func isSnuke(s [][]rune, i, j, H, W int) (bool, [][]int) {
	// up
	if i-4 >= 0 && s[i-1][j] == 'n' && s[i-2][j] == 'u' && s[i-3][j] == 'k' && s[i-4][j] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}, {i - 4, j}}
	}
	// down
	if i+4 < H && s[i+1][j] == 'n' && s[i+2][j] == 'u' && s[i+3][j] == 'k' && s[i+4][j] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}, {i + 4, j}}
	}
	// left
	if j-4 >= 0 && s[i][j-1] == 'n' && s[i][j-2] == 'u' && s[i][j-3] == 'k' && s[i][j-4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}, {i, j - 4}}
	}
	// right
	if j+4 < W && s[i][j+1] == 'n' && s[i][j+2] == 'u' && s[i][j+3] == 'k' && s[i][j+4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}, {i, j + 4}}
	}
	// left up
	if j-4 >= 0 && i-4 >= 0 && s[i-1][j-1] == 'n' && s[i-2][j-2] == 'u' && s[i-3][j-3] == 'k' && s[i-4][j-4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}, {i - 4, j - 4}}
	}
	// left down
	if j-4 >= 0 && i+4 < H && s[i+1][j-1] == 'n' && s[i+2][j-2] == 'u' && s[i+3][j-3] == 'k' && s[i+4][j-4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}, {i + 4, j - 4}}
	}
	// right up
	if j+4 < W && i-4 >= 0 && s[i-1][j+1] == 'n' && s[i-2][j+2] == 'u' && s[i-3][j+3] == 'k' && s[i-4][j+4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}, {i - 4, j + 4}}
	}
	// right down
	if j+4 < W && i+4 < H && s[i+1][j+1] == 'n' && s[i+2][j+2] == 'u' && s[i+3][j+3] == 'k' && s[i+4][j+4] == 'e' {
		i++
		j++
		return true, [][]int{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}, {i + 4, j + 4}}
	}

	return false, nil
}
