package main

import "fmt"

func main() {
	a, b := 2, 7
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
	/* обмен с помощью математических операций*/
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a: %d, b: %d\n", a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Printf("a: %d, b: %d\n", a, b)
	a = a - b
	b = a + b
	a = (-1)*a + b
	fmt.Printf("a: %d, b: %d\n", a, b)
	/* обмен с помощью исключающего или */
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("a: %d, b: %d\n", a, b)
}
