package main

import "fmt"

func fibonacci(num int) {
	a := 0
	b := 1
	c := b

	fmt.Printf("Series is %d %d", a, b)
	for true {
		c = b
		b = a + b

		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
	fmt.Println()

}

func main() {
	// fibonacci(2)
	fibonacci(3)
}
