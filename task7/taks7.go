package main

import (
	"fmt"
	"sync"
)

/* создание струтуры container со словарем и встраиванием методов RWMutex */
type Container struct {
	data map[interface{}]interface{}
	sync.RWMutex
}

/* конструктор container*/
func NewContainer() *Container {
	return &Container{data: make(map[interface{}]interface{})}
}

/* функция для конкурентной записи в словарь data структуры container*/
func (cont *Container) Write(key interface{}, value interface{}) {
	/*блокирование доступа к изменению словаря для других горутин*/
	cont.Lock()
	/*разблокирование доступа к изменению словаря для других горутин*/
	defer cont.Unlock()
	/*запись в словарь*/
	cont.data[key] = value
}

/* функция для конкурентного чтения из словаря*/
func (cont *Container) Read(key interface{}) (interface{}, bool) {
	/* чтение для нескольких горутин из словаря*/
	cont.RLock()
	/*разблокирование доступа*/
	defer cont.RUnlock()
	value, ok := cont.data[key]
	return value, ok

}
func main() {
	container := NewContainer()
	keys := []int{1, 2, 3, 4, 5}
	values := []string{"данные№1", "данные№2", "данные№3", "данные№4", "данные№5"}
	var wg sync.WaitGroup
	for i := 0; i < len(keys); i++ {
		wg.Add(1)
		go func(i int) {
			container.Write(keys[i], values[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(container.data)
}
