package main

import "fmt"

func main() {
	//Desafio 1
	for i:=1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i, "é divisível por 3")
		} else {
			fmt.Println(i, "não é divisível por 3")
		}
	}

	//Desafio 2
	for i:=1; i <= 100; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println("PinPan", i)
			} else {
				fmt.Println("Pin", i)
			}
		} else if i % 5 == 0 {
			if i % 3 == 0 {
				fmt.Println("PinPan", i)
			} else {
				fmt.Println("Pan", i)
			}
		}
	}

	//Desafio 2 - Refatorado
	for i := 1; i <= 100; i++ {
        output := ""

        if i % 3 == 0 {
            output += "Pin"
        }
        if i % 5 == 0 {
            output += "Pan"
        }

        if output != "" {
            fmt.Println(output, i)
        }
    }
}