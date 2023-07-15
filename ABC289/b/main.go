package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type connectedConpoment struct {
	min int
	max int
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var a []int
	s := strings.Split(nextLine(), " ")
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		a = append(a, i)
	}

	var p []string

	// M=0の場合はそのまま表示
	if M == 0 {
		for i := 1; i <= N; i++ {
			p = append(p, strconv.Itoa(i))
		}
		fmt.Println(strings.Join(p, " "))
		return
	}

	// 連結要素を抽出
	var cc []connectedConpoment
	min, max := a[0], a[0]
	for i := 0; i < M; i++ {
		max = max + 1

		if i+1 == M {
			cc = append(cc, connectedConpoment{min, max})
			break
		}

		if max < a[i+1] {
			cc = append(cc, connectedConpoment{min, max})
			min = a[i+1]
			max = min
			continue
		}
	}

	i := 1
	for _, c := range cc {
		for i <= N {
			if i == c.min {
				for j := c.max; j >= i; j-- {
					p = append(p, strconv.Itoa(j))
				}
				i = c.max + 1
				break
			} else {
				p = append(p, strconv.Itoa(i))
				i++
			}
		}
	}

	for i <= N {
		p = append(p, strconv.Itoa(i))
		i++
	}

	fmt.Println(strings.Join(p, " "))
}
