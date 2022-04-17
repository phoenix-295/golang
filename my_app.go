// Hello Word
package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
)

// Let us start
func main() {
	fmt.Println("Hello, world!")                          // Print simple text on screen
	fmt.Println(os.Getenv("USER"), ", Let's be friends!") // Read Linux $USER environment variable
	os.Chdir("/home/phoenix/github")
	println("Hello")
}
