package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i // Mengirim nilai i ke channel
		time.Sleep(time.Second)
	}
	close(ch) // Menutup channel setelah semua nilai dikirim
}

func consumer(ch <-chan int, done chan<- bool) {
	for {
		value, open := <-ch // Menerima nilai dari channel
		if !open {
			done <- true
			return
		}
		fmt.Println("Consuming:", value)
	}
}

func main() {
	dataChannel := make(chan int) // Membuat channel untuk data
	doneChannel := make(chan bool) // Membuat channel untuk sinyal selesai

	go producer(dataChannel) // Menjalankan goroutine produsen
	go consumer(dataChannel, doneChannel) // Menjalankan goroutine konsumen

	// Menunggu hingga goroutine konsumen selesai
	<-doneChannel
}
