package internal

import "errors"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Divide divides a by b.
// If b is zero, it returns an error.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// IsEven checks whether a number is even.
func IsEven(n int) bool {
	return n%2 == 0
}
