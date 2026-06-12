/*
Here we are going to discuss about variables in Go Lang.
- Variables are used to store data in a program. They have a name, a type, and a value.
- In Go, you can declare a variable using the var keyword followed by the variable name and type.
- You can also initialize a variable at the time of declaration.
- Go supports various data types such as int, float64, string, bool, etc.

*/

package main
import (
	"fmt"
	"flag"
)

func main() {
	fmt.Println("Hello, Go!");
	var a int = 10
	fmt.Println("Value of a:", a)

	var b string = "Go Programming"
	fmt.Println("Value of b:", b)

	var c bool = true
	fmt.Println("Value of c:", c)

	var d float64 = 3.14
	fmt.Println("Value of d:", d)

	// flag
	var name string
	flag.StringVar(&name, "name", "World", "a name to say hello to")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", name)
}