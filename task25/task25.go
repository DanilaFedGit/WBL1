package main

import (
	"fmt"
	"time"
)

func customsleep(t time.Duration) {
	/* спустя время time.Second * t time.After отправляет данные в канал*/
	<-time.After(time.Second * t)
}
func main() {
	var second time.Duration
	fmt.Scan(&second)
	start := time.Now()
	customsleep(second)
	fmt.Println(time.Since(start))

}
