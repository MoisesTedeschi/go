package main

import "fmt"

const count = 10
func main() {
	// j := 1
	// for j <= 10 {
	// 	fmt.Println(j)
	// 	j++
	// }
	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			fmt.Println("Número par:", i)
		} else {
			fmt.Println("Número ímpar:", i)
		}
	}
}