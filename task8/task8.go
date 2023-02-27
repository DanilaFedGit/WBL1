package main

import (
	"errors"
	"fmt"
)

func main() {
	var x int64 = 49
	var index1 int = 3
	fmt.Println(setone(x, index1))
	var index2 int = 5
	fmt.Println(setzero(x, index2))

}

/*функция для изменения бита под индексом index на 1*/
func setone(number int64, index int) (int64, error) {
	if index >= 0 {
		/* побитовый сдвиг 1 вправо, потом оперция или между 1 действием и number*/
		return number | (1 << index), nil
	}
	return number, errors.New("не верный индекс")
}
func setzero(number int64, index int) (int64, error) {
	if index >= 0 {
		/*побитовый сдвиг вправо, потом у результата 1 действия биты меняются на противоположный,
		в конце операцияи между результатом 2 действия и number*/
		return number & ^(1 << index), nil
	}
	return number, errors.New("не верный индекс")
}
