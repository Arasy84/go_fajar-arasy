package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func printMultiples(x int, ch chan int) {
	for i := 1; ; i++ {
		ch <- x * i
		time.Sleep(3 * time.Second)
	}
}

func main() {
	x := 7
	ch := make(chan int)
	go printMultiples(x, ch)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	timer := time.NewTimer(15 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("Program Berakhir Otomatis")
			return
		case num := <-ch:
			fmt.Println(num)
		}
	}
}

//ctrl + C buat stop program
