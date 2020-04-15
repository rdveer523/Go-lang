// Go - different types of print function
package main // package declaration

import "fmt" // importing package tells to preprocessor command to include all files from this

func main() { // func is keyword to deine the function in Go
	// main() is mandatory function which execution starts from here
	fmt.Print("Hello Mr VEERA ")      // Print() - prints but the next expression will club in this line
	fmt.Println(".....How are you.?") // Println() - prints but the next print statement will not club into this
	print("This is normal print")     // same
	print(".....clubbed")
}
