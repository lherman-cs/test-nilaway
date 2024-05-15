package main

import "fmt"

func foo() *int {
	return nil
}

func bar() {
	print(*foo()) // nilness reports NO error here, but NilAway does.
}

func hello() *int {
	return nil
}

func main() {
	t := hello()
	b := *t + 3
	fmt.Println(b)
}
