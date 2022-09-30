package main

import (
	"fmt"
	"strings"
)

func main() {
	greet("Adam", 24)
	cetak("Jl.Asalamualikum", 13456)

	var namess = []string{"ali", "eko"}
	var printMsg = greet2("Hello", namess)

	fmt.Println("-------Function(Return)-------")
	fmt.Println(printMsg)
}

func greet(name string, age int) {
	fmt.Printf("Hello nama saya %s dan umur saya %d \n", name, age)
}
func cetak(address string, number int) {
	fmt.Printf("Alamat saya di %s nomor %d \n", address, number)
}
func greet2(msg string, namess []string) string {
	var joinStr = strings.Join(namess, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}
