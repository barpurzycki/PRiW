package main

import (
	"fmt"
	"sync"
)

func main() {
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

	var kanal chan int
	kanal := make(chan int, 3)

	var wg sync.WaitGroup

	l_watkow := 2

	for i := 0; i < l_watkow; i++ {
		wg.Add(1)
	}
}
