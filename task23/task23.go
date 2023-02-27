package main

import (
	"fmt"
	"log"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var index int
	fmt.Scan(&index)
	sl = deleteindex1(sl, index)
	fmt.Println(sl)
	sl2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 = deleteindex2(sl2, index)
	fmt.Println(sl2)
}

/*удаление с сохранением порядка элементов слайса*/
func deleteindex1(sl []int, index int) []int {
	/*провекрка индекс */
	if index < 0 || index >= len(sl) {
		log.Println("Не верный индекс")
		return sl
	}
	/*удаление элемента с заданным индексом*/
	return append(sl[0:index], sl[index+1:]...)
}

/*удаление без сохранением порядка элементов слайса*/
func deleteindex2(sl []int, index int) []int {
	/*провекрка индекс */
	if index < 0 || index >= len(sl) {
		log.Println("Не верный индекс")
		return sl
	}
	/* элемент с заданным индексом меняется местами с последним элементом слайса*/
	sl[index] = sl[len(sl)-1]
	/*удаление последнего элемента слайса*/
	return sl[0 : len(sl)-1]

}
