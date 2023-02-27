package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/* создание воркеров*/
func worker(id int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range channel {
		fmt.Printf("worker №%d data: %d \n", id, i)
	}
	return

}
func main() {
	var n int
	fmt.Scan(&n)
	/* буферизированый канал для считываги ctrl+c*/
	channel_os := make(chan os.Signal, 1)
	signal.Notify(channel_os, os.Interrupt, syscall.SIGTERM)
	channel_int := make(chan int)
	var wg sync.WaitGroup
	/*создание заданного количества вооркеров*/
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, channel_int, &wg)
	}
	for {
		select {
		/* при вводе ctrl+c закрывается канал после чего оставшиеся воркеры в группе закнчивают
		работу и программа останаливается*/
		case <-channel_os:
			close(channel_int)
			fmt.Println("Прекращение работы")
			wg.Wait()
			return
		default:
			time.Sleep(2 * time.Second)
			channel_int <- rand.Int()
		}
	}

}
