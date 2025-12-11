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

func Planista(kanal_zadania chan int, kanal_ze_zwrotkami chan chan int, f func(func(int) int, int) int) {
	for x := range kanal_zadania {
		nowa := f(funkcja_1_1, x)
		kanal_na_zwrotke := <-kanal_ze_zwrotkami
		kanal_na_zwrotke <- nowa
		close(kanal_na_zwrotke)
	}
}

func main() {
	kanal_zadania := make(chan int)
	kanal_ze_zwrotkami := make(chan chan int, 1)

	go Planista(kanal_zadania, kanal_ze_zwrotkami, funkcja_sluzaca)

	fmt.Println("TEST 1")
	kanal_zadania <- 10
	kanal_zwrotki1 := make(chan int)
	kanal_ze_zwrotkami <- kanal_zwrotki1
	fmt.Println("Odpowiedź:", <-kanal_zwrotki1)

	fmt.Println("TEST 2")
	kanal_zadania <- 5
	kanal_zwrotki2 := make(chan int)
	kanal_ze_zwrotkami <- kanal_zwrotki2
	fmt.Println("Odpowiedź:", <-kanal_zwrotki2)

	close(kanal_zadania)
}
