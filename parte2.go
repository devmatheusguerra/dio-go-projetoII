package main

import "fmt"

func main() {
	for i := 2; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		}

	}
}
