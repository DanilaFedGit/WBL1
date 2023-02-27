package main

import (
	"fmt"
	"math/rand"
)

func quicksort(numbers []int) {
	/* если рзмер слайса меньше 2,то его нет смысла сортировать*/
	if len(numbers) < 2 {
		return
	}
	left := 0
	right := len(numbers) - 1
	/* index опорного элемента*/
	index := rand.Intn(right)
	/* значение опорного элемента*/
	pivot := numbers[index]
	for left <= right {
		/* пропускаем элементы, которые меньше опорного*/
		for numbers[left] < pivot {
			left += 1
		}
		/* пропускаем элементы, которые больше опорного*/
		for numbers[right] > pivot {
			right -= 1
		}
		if left <= right {
			/* элементы меньше опорного перекидываются налево, больше - направо*/
			numbers[left], numbers[right] = numbers[right], numbers[left]
			left += 1
			right -= 1

		}
	}
	/*рекурсивные вызовы ф/ии быстрой сортировки для сортировки левой и правой воловинки*/
	if right > 0 {
		quicksort(numbers[:right+1])
	}
	if left < len(numbers) {
		quicksort(numbers[left:])

	}
}
func main() {
	lst := []int{56, 34, 1, 22, 10, 123, 1234, 32, 12, 2345, 99, 1, 0, 2222222}
	quicksort(lst)
	fmt.Println(lst)
}
