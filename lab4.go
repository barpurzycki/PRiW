package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//Zadanie 1

	const size = 300
	tab := make([]float64, size)

	tab_cz1 := tab[0:100]
	tab_cz2 := tab[100:200]
	tab_cz3 := tab[200:300]

	for i := range tab_cz1 {
		tab_cz1[i] = -2.0
	}

	for i := range tab_cz2 {
		tab_cz2[i] = -4.0
	}

	for i := range tab_cz3 {
		tab_cz3[i] = 6.0
	}

	fmt.Println(tab_cz1[0])
	fmt.Println(tab_cz2[0])
	fmt.Println(tab_cz3[0])

	//Zadanie 2

	kanal2 := make(chan float64, 3)

	sumuj := func(dane []float64) {
		defer wg.Done()
		var suma float64
		for _, v := range dane {
			suma += v
		}
		kanal2 <- suma
	}

	wg.Add(3)
	go sumuj(tab_cz1)
	go sumuj(tab_cz2)
	go sumuj(tab_cz3)

	go func() {
		wg.Wait()
		close(kanal2)
	}()

	for s := range kanal2 {
		fmt.Printf("Suma czesci: %.2f\n", s)
	}
	fmt.Println()

	//Zadanie 3

	kanal3 := make(chan float64, 3)

	sumuj3 := func(dane []float64) {
		defer wg.Done()
		var suma float64
		for _, v := range dane {
			suma += v
		}
		kanal3 <- suma
	}

	pusty := func() {
		defer wg.Done()
	}

	wg.Add(8)
	go sumuj3(tab_cz1)
	go sumuj3(tab_cz2)
	go sumuj3(tab_cz3)
	for i := 0; i < 5; i++ {
		go pusty()
	}

	go func() {
		wg.Wait()
		close(kanal3)
	}()

	for s := range kanal3 {
		fmt.Printf("Suma: %.2f\n", s)
	}
	fmt.Println()

	//Zadanie 4

	kanal4 := make(chan float64, 3)

	sumuj4 := func(dane []float64) {
		defer wg.Done()
		var suma float64
		for _, v := range dane {
			suma += v
		}
		kanal4 <- suma
	}

	l_watkow := 2
	wg.Add(l_watkow)

	go sumuj4(tab_cz1)
	go sumuj4(tab_cz2)

	go func() {
		wg.Wait()
		close(kanal4)
	}()

	for s := range kanal4 {
		fmt.Printf("Suma: %.2f\n", s)
	}
	fmt.Println()

	//Zadanie 5

	kanal5 := make(chan float64, 3)

	sumuj5 := func(dane []float64) {
		defer wg.Done()
		var suma float64
		for _, v := range dane {
			suma += v
		}
		kanal5 <- suma
	}

	wg.Add(3)
	go sumuj5(tab_cz1)
	go sumuj5(tab_cz2)
	go sumuj5(tab_cz3)

	go func() {
		wg.Wait()
		close(kanal5)
	}()

	var sumaCalkowita float64
	i := 1
	for wartosc := range kanal5 {
		fmt.Printf("Suma czesci %d: %.2f\n", i, wartosc)
		sumaCalkowita += wartosc
		i++
	}

	fmt.Printf("Suma calkowita: %.2f\n", sumaCalkowita)
}
