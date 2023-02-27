package main

import (
	"fmt"
	"sync"
)

func main() {
	/* создание wait группы*/
	var wg sync.WaitGroup
	/* создание mutex*/
	var mu sync.Mutex
	numbers := []int{2, 4, 6, 8, 10}
	sumnumbers := 0
	for _, i := range numbers {
		/* добавление в wait группу*/
		wg.Add(1)
		go func(number int) {
			/* сигнал о том,что горутина закончила свое выполнение*/
			defer wg.Done()
			/* блокировка доступа для остальных гоурутин во время выполнения текущей гоурутины*/
			mu.Lock()
			/* разблокировка доступа для других гоурутин*/
			defer mu.Unlock()
			sumnumbers += number * number
		}(i)

	}
	/* ожидание выполнения горутин в wait группе*/
	wg.Wait()
	fmt.Println(sumnumbers)
}
