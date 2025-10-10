package main

import (
	"fmt"
	"math"
)

func main() {
	const n int = 1024
	var tab [n]float64
	for i := 0; i < n; i++ {
		angle := float64(i) * 2 * math.Pi / float64(n)
		tab[i] = math.Sin(angle)
	}

	var sum float64 = 0
	for i := 0; i < n; i++ {
		sum += tab[i]
	}

	fmt.Printf("%.10f", sum)
}
