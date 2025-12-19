package main

import (
	"fmt"
	"time"
)

func funkcja_1_1(x int) int {
	return x
}

func funkcja_sluzaca(f func(int) int, x int) int {
	val := f(x)
	fmt.Println("Odebrano liczbę:", val)
	time.Sleep(500 * time.Millisecond)
	val2 := val + 100
	fmt.Println("Po powiększeniu o 100:", val2)
	return val2
}

func Planista(
	kanal_zadania chan int,
	kanal_ze_zwrotkami chan chan int,
	f func(func(int) int, int) int,
) {
	fmt.Println("[Planista] Uruchomiony")

	for x := range kanal_zadania {
		fmt.Println("[Planista] Nowe zadanie:", x)

		wynik := f(funkcja_1_1, x)

		kanal_na_zwrotke := <-kanal_ze_zwrotkami
		kanal_na_zwrotke <- wynik
		close(kanal_na_zwrotke)

		fmt.Println("[Planista] Zadanie zakończone")
	}

	fmt.Println("[Planista] Zakończony")
}

type ObiektAktywny struct {
	kanal_wejsciowy      chan int
	kanal_ze_zwrotkami   chan chan int
	wybrana_funkcja      func(func(int) int, int) int
}

func (oa *ObiektAktywny) start(
	rozmiar_bufora int,
	_funkcja_sluzaca func(func(int) int, int) int,
) {
	oa.kanal_wejsciowy = make(chan int, rozmiar_bufora)
	oa.kanal_ze_zwrotkami = make(chan chan int, rozmiar_bufora)
	oa.wybrana_funkcja = _funkcja_sluzaca

	fmt.Println("[ObiektAktywny] Start – bufor:", rozmiar_bufora)

	go Planista(
		oa.kanal_wejsciowy,
		oa.kanal_ze_zwrotkami,
		oa.wybrana_funkcja,
	)
}

func (oa *ObiektAktywny) dodaj_prace_do_kanalu(liczba int) chan int {
	fmt.Println("[ObiektAktywny] Dodaję pracę:", liczba)

	kanal_na_zwrotke := make(chan int)

	oa.kanal_wejsciowy <- liczba
	oa.kanal_ze_zwrotkami <- kanal_na_zwrotke

	return kanal_na_zwrotke
}

func main() {
	var oa ObiektAktywny

	oa.start(1, funkcja_sluzaca)

	fmt.Println("\n--- TEST 1 ---")
	k1 := oa.dodaj_prace_do_kanalu(10)
	fmt.Println("Odpowiedź:", <-k1)

	fmt.Println("\n--- TEST 2 ---")
	k2 := oa.dodaj_prace_do_kanalu(5)
	fmt.Println("Odpowiedź:", <-k2)

	close(oa.kanal_wejsciowy)
}
