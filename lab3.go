package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum(tab []float64, kanalWynikow chan<- float64) {
	defer wg.Done()
	var suma float64
	for _, val := range tab {
		suma += val
	}
	fmt.Printf("Suma części: %f\n", suma)
	kanalWynikow <- suma
}

func worker(parts [][]float64, kanalCzesci <-chan int, kanalWynikow chan<- float64) {
	defer wg.Done()
	for nr := range kanalCzesci {
		tablica := parts[nr-1]
		var suma float64
		for _, v := range tablica {
			suma += v
		}
		fmt.Printf("Część %d -> suma = %f\n", nr, suma)
		kanalWynikow <- suma
	}
}

func main() {
	const size = 10000
	tab := make([]float64, size)

	tab_cz1 := tab[0:2000]
	tab_cz2 := tab[2000:5000]
	tab_cz3 := tab[5000:6500]
	tab_cz4 := tab[6500:10000]

	for i := range tab_cz1 {
		tab_cz1[i] = 1.0
	}

	for i := range tab_cz2 {
		tab_cz2[i] = 1.0 / 3.0
	}

	for i := range tab_cz3 {
		tab_cz3[i] = -1.0
	}

	for i := range tab_cz4 {
		tab_cz4[i] = -1.0 / 3.0
	}

	fmt.Println(tab_cz1[0])
	fmt.Println(tab_cz2[0])
	fmt.Println(tab_cz3[0])
	fmt.Println(tab_cz4[0])

	kanalCzesci := make(chan int, 4)
	kanalWynikow := make(chan float64, 4)

	czesci := [][]float64{tab_cz1, tab_cz2, tab_cz3, tab_cz4}

	wg.Add(2)
	go worker(czesci, kanalCzesci, kanalWynikow)
	go worker(czesci, kanalCzesci, kanalWynikow)

	for i := 1; i <= 4; i++ {
		kanalCzesci <- i
	}
	close(kanalCzesci)

	go func() {
		wg.Wait()
		close(kanalWynikow)
	}()

	var sumaCalkowita float64
	for suma := range kanalWynikow {
		sumaCalkowita += suma
	}

	fmt.Printf("Calkowita suma w kanale: %f", sumaCalkowita)

}
