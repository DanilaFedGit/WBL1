package main

import "fmt"

func main() {
	set1 := []int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}
	set2 := []int{11, 12, 1, 2, 689, 6}
	fmt.Println(SetIntersection(set1, set2))

}
func SetIntersection(set1 []int, set2 []int) []int {
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	/* слайс пересечений*/
	intersection := make([]int, 0, 0)
	/*словарь для значений из первого мнжества*/
	setmap := make(map[int]bool)
	for _, i := range set1 {
		setmap[i] = true
	}
	for _, i := range set2 {
		/*если значение из второго множества есть в словаре, то оно есть и в первом*/
		if _, ok := setmap[i]; ok {
			intersection = append(intersection, i)
		}
	}
	/* слайс пересечений двух множеств*/
	return intersection

}
