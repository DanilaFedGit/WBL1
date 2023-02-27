package main

import "fmt"

/*структура Human*/
type Human struct {
	name string
	age  int
}

/*метод структуры Human*/
func (person *Human) Getdata() {
	fmt.Printf("имя: %s, возраст: %d\n", person.name, person.age)
}

type Action struct {
	/* встраивание Human в структуру Action*/
	Human
}

func main() {
	act := Action{Human{"Bob", 19}}
	act.Getdata()
}
