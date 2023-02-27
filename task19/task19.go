package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	fmt.Println(reverstring(scaner.Text()))

}
func reverstring(str string) string {
	/* создание массива рун*/
	runes := []rune(str)
	for i, k := 0, len(runes)-1; i < k; i, k = i+1, k-1 {
		/* меняет места i элемент c k элементом*/
		runes[i], runes[k] = runes[k], runes[i]
	}
	/* преобразование рун в строку*/
	return string(runes)
}
