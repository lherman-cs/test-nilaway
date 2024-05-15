package main

import "fmt"

func hello() *int {
	return nil
}

func main() {
	t := hello()
	b := *t + 3
	fmt.Println(b)
}
