package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	height, width float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var cl shape = circle{radius: 5}
	var r1 shape = rectangle{width: 10, height: 5}

	fmt.Printf("Type of c1: %T ", cl)
	fmt.Printf("Type of r1: %T", r1)

	// user.StoreUser(30, "Joe")

	// fmt.Println(user.GetUsers(4))

	Data()
}

func Data() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println(" type of variabel:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("value of variabel:", reflectValue.Int())
	}
}
