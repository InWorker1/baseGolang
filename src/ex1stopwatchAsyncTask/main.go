package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"sort"
	"sync"
	"time"
)

type info struct {
	number   int
	duration int
}

func routine(timeSleep int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(timeSleep))
}

func main() {
	nptr, mptr := flag.Int("N", 0, ""), flag.Int("M", 0, "")
	flag.Parse()
	n, m := *nptr, *mptr
	var wg sync.WaitGroup

	list := make([]info, 0, n)
	for i := 1; i < n+1; i++ {
		if m <= 0 {
			fmt.Println("M не может быть отрицательным")
			return
		}
		timeSleep := rand.IntN(m + 1)
		wg.Add(1)
		go routine(timeSleep, &wg)
		list = append(list, info{number: i, duration: timeSleep})
	}
	wg.Wait()

	sort.Slice(list, func(i, j int) bool {
		return list[i].duration > list[j].duration
	})

	for _, v := range list {
		fmt.Printf("%d, %d\n", v.number, v.duration)
	}
}
