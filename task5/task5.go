package main

import (
	"fmt"
	"time"
)

func main() {
	/* время старта программы*/
	/* создание канала*/
	channel := make(chan int)
	/* запуск горутины*/
	go writedata(channel)
	/* переменная для ввода количества секунд работы программы*/
	var durationtime time.Duration
	fmt.Scan(&durationtime)
	/* запуск таймера*/
	to := time.After(durationtime * time.Second)
	/* флаг для цикла*/
	var timeend bool = true
	for timeend {
		select {
		/* если таймер прошел, то флаг становится ложным*/
		case <-to:
			timeend = false
		/* чтение данных из канала*/
		case data := <-channel:
			fmt.Println(data)
		}
	}
	close(channel)

}
func writedata(channel chan int) {
	/* отправка данных в канал*/
	for i := 0; ; i++ {
		channel <- i
	}
}
