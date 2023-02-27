package main

import (
	"fmt"
)

func gen(done chan struct{}, numbers []int) <-chan int {
	/*канал для записи чисел из массива*/
	out := make(chan int)
	go func() {
		/*закрытие канала после выполнения функции*/
		defer close(out)
		/* как только прекращается main прекращаем отправлять данные*/
		for _, i := range numbers {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}
func sq(done chan struct{}, in <-chan int) <-chan int {
	/* канал для записи квадратов чисел из канала in*/
	out := make(chan int)
	go func() {
		/*закрытие канала после выполнения функции*/
		defer close(out)
		/* как только прекращается main прекращаем отправлять данные*/
		for i := range in {
			select {
			case out <- i * i:
			case <-done:
				return
			}
		}
	}()
	return out
}
func main() {
	done := make(chan struct{})
	defer close(done)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 9}
	/*вывод данных из второго канала*/
	for i := range sq(done, gen(done, numbers)) {
		fmt.Println(i)
	}

}
