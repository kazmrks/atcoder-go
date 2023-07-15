package main

import "fmt"

func main() {
	var S, out string
	fmt.Scan(&S)
	for _, c := range S {
		if c == '0' {
			out += "1"
		} else if c == '1' {
			out += "0"
		}
	}
	fmt.Println(out)
}
