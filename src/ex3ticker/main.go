package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ticker(k uint, ctx context.Context, chanExitRoutine chan struct{}) {
	i := 1
	counter := 0
	defer close(chanExitRoutine)
	for {
		fmt.Printf("Tick %d since %d\n", i, counter)
		time.Sleep(time.Second * time.Duration(k))
		i++
		counter += int(k)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	kptr := flag.Uint("K", 0, "K")
	flag.Parse()
	k := *kptr

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	chanExitRoutine := make(chan struct{})

	go ticker(k, ctx, chanExitRoutine)

	<-osChan
	cancel()
	<-chanExitRoutine
	fmt.Println("Termination")
}
