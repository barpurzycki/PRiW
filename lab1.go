package main

import (
	"fmt"
	"math"
	"time"
)

const n int = 1024

var threads int = 16
var tab [n]float64
var sum2 float64 = 0

func suma(id int, segment int) {
	start := id * segment
	end := start + segment
	localSum := 0.0
	for i := start; i < end; i++ {
		localSum += tab[i]
	}
	sum2 += localSum
}

func main() {
	// Zadanie 1 i 2

	for i := 0; i < n; i++ {
		angle := float64(i) * 2 * math.Pi / float64(n)
		tab[i] = math.Sin(angle)
	}

	var sum float64 = 0
	for i := 0; i < n; i++ {
		sum += tab[i]
	}

	fmt.Printf("Suma 1 wątku: %.20f\n", sum)

	// Zadanie 3

	segment := n / threads
	for j := 0; j < threads; j++ {
		go suma(j, segment)
	}

	time.Sleep(1 * time.Second)

	fmt.Printf("Suma 16 wątków: %.20f", sum2)

}
