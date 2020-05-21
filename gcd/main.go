package main

import "fmt"

func findGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return findGcd(b, a%b)
}

func main() {
	var a, b = 1701, 3768

	fmt.Println(findGcd(a, b))
}
