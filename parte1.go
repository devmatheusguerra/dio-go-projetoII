package main

import "fmt"

func main() {
	for i := 2; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
