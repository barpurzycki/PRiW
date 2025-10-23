package main

import (
	"fmt"
	"sync"
)

var mut sync.RWMutex
var wg sync.WaitGroup

const n int = 10000

var tab [n]float64

func statystyka(tab *[n]float64) (float64, float64) {
	sum := 0.0
	max := tab[0]
	for _, val := range tab {
		sum += val
		if val > max {
			max = val
		}
	}
	return sum, max
}

func czytacz() {
	defer wg.Done()
	mut.RLock()
	sum, max := statystyka(&tab)
	fmt.Printf("Suma: %.2f, Max: %.2f\n", sum, max)
	mut.RUnlock()
}

func pisarzMnoznik() {
	defer wg.Done()
	mut.Lock()
	for i := range tab {
		tab[i] *= 2
	}
	mut.Unlock()
}

func main() {

	for i := 0; i < n/2; i++ {
		tab[i] = 1
	}

	for i := n / 2; i < n; i++ {
		tab[i] = -1
	}

	fmt.Printf("Zadanie 2:\n")
	sum, max := statystyka(&tab)
	fmt.Printf("Suma wartości tablicy: %.2f, Maksymalna wartość: %.2f\n\n", sum, max)

	fmt.Printf("Zadanie 3:\n")
	wg.Add(3)
	go czytacz()
	go czytacz()
	go czytacz()
	wg.Wait()

	sum, max = statystyka(&tab)
	fmt.Printf("Suma wartości tablicy: %.2f, Maksymalna wartość: %.2f\n\n", sum, max)

	fmt.Printf("Zadanie 4:\n")
	wg.Add(3)
	go pisarzMnoznik()
	go pisarzMnoznik()
	go pisarzMnoznik()
	wg.Wait()

	sum, max = statystyka(&tab)
	fmt.Printf("Suma wartości tablicy: %.2f, Maksymalna wartość: %.2f\n\n", sum, max)

	fmt.Printf("Zadanie 5:\n")
	wg.Add(5)
	go czytacz()
	go czytacz()
	go czytacz()
	go pisarzMnoznik()
	go pisarzMnoznik()
	wg.Wait()

	sum, max = statystyka(&tab)
	fmt.Printf("Suma wartości tablicy: %.2f, Maksymalna wartość: %.2f\n", sum, max)
}
