package main

import "fmt"

func main() {
	var t *int
	b := *t + 3
	fmt.Println(b)
}
