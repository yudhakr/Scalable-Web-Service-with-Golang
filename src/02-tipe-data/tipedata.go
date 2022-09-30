package main

import "fmt"

func main() {

	/*imutables
	int -> unit,unit8,unit16,unit12 . . . int64
	float -> 32 dan 64
	string
	bool
	*/

	var namadepan string = "rehan"
	var namablk string = "surya"

	fmt.Printf("halo %s %s \n", namadepan, namablk)

	// multi declaration
	a1, a2, a3, a4 := "string1", "string2", "string3", "string4"
	fmt.Println(a1, a2, a3, a4)

	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first %T \n", first)
	fmt.Printf("tipe data second %T \n", second)
	//tipe data Number(interger)
	//tipe data float
	//Tipe data number(float)

	var decimalnumber float32 = 3.64

	fmt.Printf("decimal Number:  %f\n", decimalnumber)
	fmt.Printf("decimal Number:  %.2f\n", decimalnumber)
	fmt.Printf("decimal Number:  %.3f\n", decimalnumber)

	//tipe data bool

	var kondisi bool = true
	fmt.Printf("kondisi %t \n", kondisi)

	// Operator
	var value = 2 + 2*3
	fmt.Println(value)
}
func Constant() {

	const full_name string = "Rehan Surya"

	fmt.Printf("Hello %s", full_name)
}
