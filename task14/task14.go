package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number interface{} = 1
	var str interface{} = "asgga"
	var t interface{} = false
	var sliceofnumber interface{} = []int{1, 2, 3}
	fmt.Println(checktype(number))
	fmt.Println(checktype(str))
	fmt.Println(checktype(t))
	fmt.Println(checktype(sliceofnumber))
}

/* функция проверки типа интерфейса*/
func checktype(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}
