package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	/* создание из введеной строки переменной типа big.Int в 10-ной системе*/
	numberA, _ := new(big.Int).SetString(a, 10)
	numberB, _ := new(big.Int).SetString(b, 10)
	/*сумма*/
	var sum big.Int
	sum.Add(numberA, numberB)
	fmt.Println("сумма: ", sum.Text(10))
	/*разность*/
	var difference big.Int
	difference.Sub(numberA, numberB)
	fmt.Println("разница: ", difference.Text(10))
	/*произведение*/
	var mul big.Int
	mul.Mul(numberA, numberB)
	fmt.Println("произведение: ", mul.Text(10))
	/*целочистленное деление*/
	var div big.Int
	div.Div(numberA, numberB)
	fmt.Println("частное: ", div.Text(10))

}
