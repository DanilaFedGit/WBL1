package main

import (
	"fmt"
	"sync"
)

/* структура Container*/
type Container struct {
	counter int
	sync.Mutex
}

/* Функция для увеличения счетчика*/
func (container *Container) Inccounter() {
	/* закрытие доступа для других горутин*/
	container.Lock()
	/* открытие доступа для других горутин*/
	defer container.Unlock()
	/*увеличение счетчика*/
	container.counter += 1

}

/* возврат значения счетчика*/
func (container *Container) Getcounter() int {
	return container.counter
}
func main() {
	/*создание Wait группы*/
	var wt sync.WaitGroup
	container := Container{}
	for i := 0; i < 100; i++ {
		/*добавление в Wait группу*/
		wt.Add(1)
		go func() {
			container.Inccounter()
			/* выход из Wait группы*/
			wt.Done()
		}()
	}
	/*ожидание выполнения горутин*/
	wt.Wait()
	fmt.Println(container.Getcounter())

}
