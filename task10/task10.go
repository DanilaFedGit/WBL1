package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.2}
	/* словарь температур с шагом 10*/
	mapoftemperatures := make(map[int][]float32)
	for _, i := range temperatures {
		/* нахождения температуры группы*/
		temperature := int(i) - int(i)%10
		/*если данной группы нет в словаре, то сохдаем эту группу*/
		if _, ok := mapoftemperatures[temperature]; !ok {
			mapoftemperatures[temperature] = make([]float32, 0, 0)
			mapoftemperatures[temperature] = append(mapoftemperatures[temperature], i)
		} else {
			/*добавление в нужную группу*/
			mapoftemperatures[temperature] = append(mapoftemperatures[temperature], i)
		}

	}
	fmt.Println(mapoftemperatures)
}
