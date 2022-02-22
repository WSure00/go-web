package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("%d is double", i)
		} else if i%5 == 0 {
			fmt.Println("%d is five", i)
		} else if i%2 != 0 {
			fmt.Println("%d is even", i)
		} else {
			fmt.Println("%d is normal", i)
		}

	}
}
