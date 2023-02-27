package main

import (
	"fmt"
	"math/rand"
)

func createHugeString(n int) []byte {
	var str string = "abcdefghijklmnopqrstuvwxyz0123456789"
	sliseofbyte := make([]byte, n, n)
	for i := 0; i < n; i++ {
		sliseofbyte[i] = str[rand.Intn(len(str))]
	}
	return sliseofbyte

}
func someFunc(n int) string {
	sliseofbyte := createHugeString(n)
	justbytes := make([]byte, 100)
	/* копирование необходимого количества символов*/
	copy(justbytes, sliseofbyte[:100])
	return string(justbytes)
}
func main() {
	fmt.Println(someFunc((1 << 10)))

}
