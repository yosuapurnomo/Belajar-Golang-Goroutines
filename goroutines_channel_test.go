package Belajar_Golang_Goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Cara untuk dapat mengembalikan data adalah dengan cara menggunakan channel yang hanya bisa menampung 1 data
func TestChannel(t *testing.T) {
	// setup channel dengan menambahkan 1 tipe datanya
	channel := make(chan string)

	go func() {
		// time.Sleep(2 * time.Second)
		// Mengirim Data sesuai tipe data yang telah diset diatas
		channel <- "Yosua"
		fmt.Println("Selesai Mengirim Data")
	}()

	// Mengambil Data yang sebelumnya telah dikirim
	data := <-channel
	// fmt.Println(<-channel)
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
