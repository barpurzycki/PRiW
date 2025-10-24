package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum(tab []float64, wg *sync.WaitGroup, kanal chan<- float64) {
	defer wg.Done()
	var suma float64
	for _, val := range tab {
		suma += val
	}
	fmt.Printf("Suma: %f\n", suma)
	kanal <- suma
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

	kanal := make(chan float64, 4)

	wg.Add(4)
	go sum(tab_cz1, &wg, kanal)
	go sum(tab_cz2, &wg, kanal)
	go sum(tab_cz3, &wg, kanal)
	go sum(tab_cz4, &wg, kanal)
	wg.Wait()
	close(kanal)

	var sumaCalkowita float64 = 0.0
	for suma := range kanal {
		sumaCalkowita += suma
	}

	fmt.Printf("Calkowita suma w kanale: %f", sumaCalkowita)

}
