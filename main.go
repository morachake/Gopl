package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Moracha")
	fmt.Println(os.Getenv("USER"))
}