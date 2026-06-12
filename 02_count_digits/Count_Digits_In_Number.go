package main

import "fmt"

func main() {
	var num, count int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num == 0 {
		count = 1
	} else {
		for num != 0 {
			count++
			num /= 10
		}
	}

	fmt.Println("Number of digits:", count)
}