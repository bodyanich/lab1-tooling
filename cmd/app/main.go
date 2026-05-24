package main

import (
	"fmt"
	"log"

	"github.com/bodyanich/lab1-tooling/internal"
)

func main() {
	sum := internal.Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", sum)

	result, err := internal.Divide(10, 2)
	if err != nil {
		log.Fatalf("failed to divide numbers: %v", err)
	}

	fmt.Printf("10 / 2 = %d\n", result)

	number := 4
	fmt.Printf("Is %d even? %v\n", number, internal.IsEven(number))
}
