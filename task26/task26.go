package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkuniqueletters(str))
}
func checkuniqueletters(str string) bool {
	/*создание словаря для проверки уникальности символов*/
	mapofletters := make(map[string]bool)
	for _, i := range str {
		/* если символ до этого встречался, то функция возвращает false*/
		if mapofletters[strings.ToUpper(string(i))] {
			return false
		}
		mapofletters[strings.ToUpper(string(i))] = true
	}
	/*возвращает true, если строка уникальна*/
	return true
}
