package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			for k := 0; k <= 999; k++ {
				go fmt.Printf("I%03d-J%03d-K%03d %s\n", i, j, k, uuid.New())
			}
		}
	}
}
