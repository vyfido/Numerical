package main

import (
	"./number"
	"fmt"
)

func main() {
	digits := number.DescomposePrime(100)
	for i, v := range digits {
		fmt.Println(i, v)
	}
	fmt.Printf("Hola mundo con Go %t\n", number.IsPerfect(12))

}
