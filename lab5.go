package main

import (
	"fmt"
	"time"
)

func l_pierwsza(kanal chan int) {
	n := 2

	for {
		czyPierwsza := true

		if n < 2 {
			czyPierwsza = false
		} else {
			for i := 2; i < n; i++ {
				if n%i == 0 {
					czyPierwsza = false
					break
				}
			}
		}
		if czyPierwsza {
			kanal <- n
			time.Sleep(1 * time.Millisecond)
		}

		n++
	}
}

func main() {
	var kanal1, kanal2, kanal3 chan int
	kanal1 = make(chan int)
	kanal2 = make(chan int)
	kanal3 = make(chan int)

	go l_pierwsza(kanal1)
	go l_pierwsza(kanal2)
	go l_pierwsza(kanal3)

	select {
		case liczba_pierwsza := <- kanal1:
			fmt.Println("Odczyt z kanalu 1, liczba pierwsza: ", liczba_pierwsza)
		case liczba_pierwsza := <- kanal2:
			fmt.Println("Odczyt z kanalu 2, liczba pierwsza: ", liczba_pierwsza)
		case liczba_pierwsza := <- kanal3:
			fmt.Println("Odczyt z kanalu 3, liczba pierwsza: ", liczba_pierwsza)
	}
}
