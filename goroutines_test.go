package Belajar_Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Goroutines tidak dapat menggembalikan data
func SayHello() {
	fmt.Println("Hello World")
}

func TestSayHello(t *testing.T) {
	// Cara Menjalankan Goroutines
	go SayHello()
	fmt.Println("Selesai")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number: ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

// Cara untuk dapat mengembalikan data adalah dengan cara menggunakan channel yang hanya bisa menampung 1 data
