package main

import (
	"fmt"

	"example.com/go-lint-demo/mathutil"
)

func main() {
	fmt.Println("2 + 3 =", mathutil.Add(2, 3))
	fmt.Println("10 / 2 =", mathutil.Divide(10, 2))
}
