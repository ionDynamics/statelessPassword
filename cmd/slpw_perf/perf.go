package main

import (
	"fmt"
	"time"

	"go.iondynamics.net/statelessPassword"
)

func main() {
	max := 8
	fmt.Println("Testing performance of", max, "variants")
	fmt.Println("\nThis may take a minute...\n\n\n")
	for i := 0; i < max; i++ {
		fmt.Println("Test", i+1, "of", max, "begins")
		start := time.Now()
		algo, _ := statelessPassword.New([]byte("Example Full Name"), []byte("example password"), uint8(i))
		algo.Password("example.org", "1", statelessPassword.Printable32Templates)
		fmt.Println(time.Since(start), "\t", "Variant code:", i, "\n")
	}
}
