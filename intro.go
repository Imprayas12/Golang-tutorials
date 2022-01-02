/* Hi, This is gonna be the first program you're looking at if you want to start programming in Go.

But before you start coding in Go. I am assuming that you are already familiar with some
programming languages. So let's create our first program
*/

package main // Just like in Java there's a concept of saving classes and methods inside packages
// Same goes with Golang. main package contains our main method

import "fmt" // Fmt is a package that contains various functions to be used including input-output
// functions same as stdio.h in C and iostream in C++

func main() { // Declaration Of main function
	fmt.Print("Hello World normal print function ")
	fmt.Println()                                            // to  change output lines
	fmt.Printf("Hello world using formatted print function") // resemblense with printf in c, c++ and java
	fmt.Println()
	fmt.Println("Hello world using new line print function ") // resemblense with println in Java
}
