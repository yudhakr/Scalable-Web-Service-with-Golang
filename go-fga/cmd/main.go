package main

import (
	"fmt"
	"os"
	// "github.com/Calmantara/go-fga/user"
)

func main() {

	arg_len := len(os.Args[1:])
	fmt.Println("Argument length: ", arg_len)

}
