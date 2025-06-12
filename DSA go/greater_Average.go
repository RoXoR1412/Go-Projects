package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if (a+b)/2 > c {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
