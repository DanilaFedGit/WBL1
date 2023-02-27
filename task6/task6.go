package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/* остановка выполнения горутины через закрытие канала*/
	channel := make(chan int)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for {
			value, ok := <-channel
			fmt.Println(value, ok)
			if !ok {
				fmt.Println("канал закрыт")
				defer wg.Done()
				return
			}
		}
	}()
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()
	channel1 := make(chan int)
	exit := make(chan bool)
	/* остановка выполнения горутины через отправку в канал данных для закрытия*/
	go func() {
		for {
			select {
			case n := <-channel1:
				fmt.Println(n)
			case <-exit:
				fmt.Println("канал закрыт")
				return
			}
		}

	}()
	channel1 <- 1
	channel1 <- 2
	exit <- true
	wg.Add(1)
	go func() {
		t := time.After(4 * time.Nanosecond)
		i := 0
		for {
			select {
			case <-t:
				defer wg.Done()
				fmt.Println("канал закрыт")
				return
			default:
				i += 1
				fmt.Println(i)
			}
		}

	}()
	wg.Wait()

}
