package main

import "fmt"

func main() {

	var name string
	var age int = 23

	name = "asep"
	age = 20
	name2 := "alibaba"
	age2 := 24

	first, secound, third := "1", 2, "3"

	fmt.Println("nama saya ==>", name)
	fmt.Println("umur saya ==>", age)

	_, _, _ = name2, age2, first

	// fmt.Println("%T, %T", name2, age2)
	fmt.Println(first, secound, third)

	// mengatahui tipe data
	fmt.Printf("Tipe data adalah %T", secound)
}
