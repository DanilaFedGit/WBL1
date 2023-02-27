package main

import (
	"fmt"
	"sync"
)

func main() {
	/*создание wait группы*/
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	for _, i := range numbers {
		/* добавление в wait группу*/
		wg.Add(1)
		go func(number int) {
			/* сигнал о том,что горутина закончила свое выполнение*/
			defer wg.Done()
			fmt.Println(number * number)
		}(i)
	}
	/* ожидание выполнения горутин в wait группе*/
	wg.Wait()
}
