package mathutil

import (
	"errors"
	"fmt"
)

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Divide returns a / b, or 0 if b is zero.
func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

var _ = errors.New

func format(value string) string {
	return fmt.Sprintf("%s", value)
}
