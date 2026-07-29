package main

import (
	"flag"
	"fmt"
)

func generator(k, n int) <-chan int {
	channel1 := make(chan int)
	go func() {
		defer close(channel1)
		for i := k; i <= n; i++ {
			channel1 <- i
		}
	}()
	return channel1
}

func squrt(channel1 <-chan int) <-chan int {
	channel2 := make(chan int)
	go func() {
		defer close(channel2)
		for {
			val, ok := <-channel1
			if !ok {
				break // канал 1 закрыт
			}
			channel2 <- val * val
		}
	}()
	return channel2
}

func main() {
	kptr, nptr := flag.Int("K", 0, ""), flag.Int("N", 0, "")
	flag.Parse()
	k, n := *kptr, *nptr

	channel1 := generator(k, n)
	channel2 := squrt(channel1)

	for i := 0; i < n-k+1; i++ {
		val, ok := <-channel2
		if ok {
			fmt.Println(val)
		}
	}

}
