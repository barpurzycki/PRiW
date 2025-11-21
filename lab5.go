package main

import (
	"fmt"
	"time"
)

func liczPierwsze(kanalCzasNaKoniec chan bool) {
	n := 2
	ostatniaPierwsza := 2

	for {
		select {
		case <-kanalCzasNaKoniec:
			fmt.Println("Koniec liczenia. Ostatnia liczba pierwsza:", ostatniaPierwsza)
			return

		default:
			czyPierwsza := true

			for i := 2; i < n; i++ {
				if n%i == 0 {
					czyPierwsza = false
					break
				}
			}

			if czyPierwsza {
				ostatniaPierwsza = n
				time.Sleep(1 * time.Millisecond)
			}

			n++
		}
	}
}

func main() {
	kanalCzasNaKoniec := make(chan bool)

	go liczPierwsze(kanalCzasNaKoniec)

	time.Sleep(10 * time.Millisecond)
	kanalCzasNaKoniec <- true

	time.Sleep(1 * time.Second)
}
