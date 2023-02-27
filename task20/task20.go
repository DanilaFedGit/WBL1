package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	fmt.Println(reverseword(scaner.Text()))
}

/* ф/ия для переворачивания слов в строке*/
func reverseword(str string) string {
	/* слайс слов из строки */
	arrayofwords := strings.Split(str, " ")
	/* цикл для переворачивания элементов в слайсе*/
	for i, k := 0, len(arrayofwords)-1; i < k; i, k = i+1, k-1 {
		arrayofwords[i], arrayofwords[k] = arrayofwords[k], arrayofwords[i]
	}
	/* элементы перевернутого слайса соединяются в строчку*/
	return strings.Join(arrayofwords, " ")
}
