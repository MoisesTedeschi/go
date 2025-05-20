package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	source := rand.NewSource(time.Now().UnixNano())
	num := rand.New(source)
	return num.Intn(10)
}
func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!", i)
	} else {
		fmt.Println("Perdeu!!!", i)
	}
}