package main

import "fmt"

func Jalan() {
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//Array
	var numbers2 [4]int
	numbers2 = [4]int{1, 5, 8, 4}

	var strings = [3]string{"asep", "lili", "budi"}

	fmt.Printf("%#v \n", numbers2)
	fmt.Printf("%#v \n", strings)
}
