package main

import (
	"fmt"
	"sync"
	//"time"
)

var mut sync.RWMutex
var wg sync.WaitGroup

const n int = 10000

var tab [n]float64

func statystyka(tab [n]float64) float64 {
	sum := 0.0
	for _, val := range tab {
		sum += val
	}
	return sum
}

func czytacz() {
	defer wg.Done()
	mut.RLock()
	fmt.Println(statystyka(tab))
	mut.RUnlock()
}

func main() {

	for i := 0; i < n/2; i++ {
		tab[i] = 1
	}

	for i := n / 2; i < n; i++ {
		tab[i] = -1
	}

	fmt.Println("Suma wartości tablicy: ", statystyka(tab))
	wg.Add(3)
	go czytacz()
	go czytacz()
	go czytacz()
	wg.Wait()
}
