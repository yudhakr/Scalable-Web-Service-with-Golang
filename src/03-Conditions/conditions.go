package main

import "fmt"

func main() {
	currentYear := 2022

	if age := currentYear - 1998; age < 18 {
		fmt.Println("belum dewasa")
	} else {
		fmt.Println("sudah dewasa")
	}

	//Switch
	//Nest Conditions

	var score2 = 0

	if score2 > 7 {
		switch score2 {
		case 10:
			fmt.Println("Sangat Bagus")
		default:
			fmt.Println("Bagus!!")
		}
	} else {
		if score2 == 5 {
			fmt.Println("Coba lagi")
		} else if score2 == 3 {
			fmt.Println("kurang")
		} else {
			fmt.Println("Jangan menyerah")
			if score2 == 0 {
				fmt.Println("Anda harus belajar lebih giat")
			}

		}
	}
}
