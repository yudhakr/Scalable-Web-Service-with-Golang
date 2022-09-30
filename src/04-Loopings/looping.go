package main

import "fmt"

func main() {

	// Loopings
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke", i)
	}
	//Lopping (Label)

	for j := 0; j < 5; j++ {
		if j == 3 {
			break
		}
		fmt.Println(j, " ")
	}
	fmt.Printf("\n")
	//Array
	var numbers2 [4]int
	numbers2 = [4]int{1, 5, 8, 4}

	var strings = [3]string{"asep", "lili", "budi"}
	strings[2] = "rehan"

	fmt.Printf("%#v \n", numbers2)
	fmt.Printf("%#v \n", strings)
	fmt.Printf("%#v \n", strings[2])

	//Array (multi dimensi)

	balances := [3][2]int{{2, 4}, {2, 1}, {3, 4}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	//Slice (append function with ellipsis)

	var buah = []string{"apel", "jeruk", "mangga"}
	var buah2 = []string{"melon", "semangka", "anggur"}

	buah = append(buah, buah2...)

	//  Slice copy
	nn := copy(buah, buah2)
	fmt.Printf("%#v", buah)

	fmt.Println("Copy element =>", nn)

	// Penggunaan Map
	var person map[string]string //Deklarasi

	person = map[string]string{} //Inisialisasi

	person["name"] = "Akbar Hermanto"

	person["age"] = "18"

	person["address"] = "Jl. Cipto Mangunkusumo No. 1"

	fmt.Println("----- Map 1-------")
	fmt.Println("nama:", person["name"])
	fmt.Println("usia:", person["age"])
	fmt.Println("alamat:", person["address"])

	// Penggunaan Map (2)
	var person2 = map[string]string{
		"nama":   "Budi Widodo",
		"usia":   "20",
		"alamat": "Jl. Anggur No. 2",
	}
	fmt.Println("------ Map 2------")
	fmt.Println("nama dia:", person2["nama"])
	fmt.Println("usia dia:", person2["usia"])
	fmt.Println("alamat dia:", person2["alamat"])

	city := "balikpapan"

	for i := 0; i < len(city); i++ {
		fmt.Printf("------ Silice------")
		fmt.Printf("index %d :byte: %d \n", i, city[i])
	}
}
