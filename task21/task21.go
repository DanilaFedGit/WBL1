package main

import "fmt"

/* interface, под который нужно адаптировать струтуру Phone*/
type Coordinates interface {
	GetCoordinates()
}

/* структура, которая соответствует заданному interface Coordinates*/
type GPSNavigator struct {
	x int
	y int
	z int
}

func (gps *GPSNavigator) GetCoordinates() {
	fmt.Println("Коодинаты: ", gps.x, gps.y, gps.z)
}

/* струтура, которая должна быть адаптирована под Coordinates*/
type Phone struct {
}

func (phone *Phone) GetLocation() {
	fmt.Println("Местоположение найдено")
}

/* адаптер для структуры Phone*/
type phoneAdapter struct {
	phone *Phone
}

/* реализация метода GetCoordinates для адаптации Phone под interface Coordinates*/
func (phoneGps *phoneAdapter) GetCoordinates() {
	phoneGps.phone.GetLocation()
}
func main() {
	var coordinate Coordinates = &GPSNavigator{1, 1, 1}
	coordinate.GetCoordinates()
	coordinate = &phoneAdapter{&Phone{}}
	coordinate.GetCoordinates()

}
