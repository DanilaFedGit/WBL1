package main

import "fmt"

func main() {
	var stringofstr = []string{"cat", "dog", "cat", "cat", "tree"}
	set := make(map[string]struct{})
	/* создание множества*/
	for _, i := range stringofstr {
		set[i] = struct{}{}
	}
	fmt.Println(set)
}
