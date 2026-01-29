package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex
var konto int = 1000
var zablokowane int = 0


func status(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	dostepne := konto - zablokowane
	wToku := "nie"
	if zablokowane > 0 {
		wToku = "tak"
	}

	fmt.Fprintf(w, "Status\nkonto: %d\nzablokowane: %d\ndostepne: %d\ntransakcje w toku: %s", konto, zablokowane, dostepne, wToku)
}

func pobierz(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/pobierz/")
	rest = strings.Trim(rest, "/")

	if rest == "" {
		http.Error(w, "Podaj kwotę w odpowiedni sposób: /pobierz/100", http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(rest)
	if err != nil || x <= 0 {
		http.Error(w, "Kwota musi być liczbą całkowitą dodatnią: /pobierz/200", http.StatusBadRequest)
		return
	}

	log.Printf("Próba pobrania, start, kwota: %d", x)

	delay := time.Duration(100+rand.Intn(201)) * time.Millisecond
	time.Sleep(delay)

	mu.Lock()
	dostepne := konto - zablokowane

	if dostepne < x {
		y := dostepne
		z := zablokowane
		mu.Unlock()

		wiadomosc := fmt.Sprintf("Próba pobrania, blokada niemożliwa, kwota: %d, pozostało niezablokowanych środków: %d i zablokowanych: %d\n", x, y, z)

		log.Print(strings.TrimSpace(wiadomosc))
		fmt.Fprint(w, wiadomosc)

		return
	}

	zablokowane += x
	mu.Unlock()

	log.Printf("Próba pobrania, blokada kwoty: %d", x)

	time.Sleep(3 * time.Second)

	mu.Lock()
	konto -= x
	zablokowane -= x
	y := konto - zablokowane
	z := zablokowane
	mu.Unlock()

	wiadomosc := fmt.Sprintf("Próba pobrania, sukces, kwota: %d, pozostało niezablokowanych środków: %d i zablokowanych: %d\n", x, y, z)
	log.Print(strings.TrimSpace(wiadomosc))
	fmt.Fprint(w, wiadomosc)
}

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/pobierz/", pobierz)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK, path=%s\n", r.URL.Path)
	})

	log.Println("Start: http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
