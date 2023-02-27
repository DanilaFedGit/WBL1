package main

import "fmt"

func main() {
	sliceofintL := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarysearch(sliceofintL, 11))
	fmt.Println(binarysearch(sliceofintL, 6))

}

/* функция бинарного поиска*/
func binarysearch(sliceofdata []int, target int) int {
	/* индекс элемента с левого конца*/
	left := 0
	/* индекс элемента с правого конца*/
	right := len(sliceofdata) - 1
	for left <= right {
		/* индекс среднего элемента между left и right*/
		middle := (left + right) / 2
		/* проверка на нахождения нужного элемента*/
		if target == sliceofdata[middle] {
			return middle
			/* если искомый элемент меньше чем средний, то сдвигает правую границу до middle-1*/
		} else if target < sliceofdata[middle] {
			right = middle - 1
			/* если искомый элемент больше чем средний, то сдвигает левую границу до middle+1*/
		} else {
			left = middle + 1
		}
	}
	return -1

}
