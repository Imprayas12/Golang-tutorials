// This program will implement the use of various identifiers,datatypes,variables and keywords which
//are available in Go programming language so let's GOPHER it.

/* List of predeclared identifiers in Go
For Constants:
true, false, iota, nil

For Types:
int, int8, int16, int32, int64, uint,
uint8, uint16, uint32, uint64, uintptr,
float32, float64, complex128, complex64,
bool, byte, rune, string, error

For Functions:
make, len, cap, new, append, copy, close,
delete, complex, real, imag, panic, recover

DataTypes :- In Go, we have 4 types of Datatypes that are Basic, Aggregate, Refrence and Interface
Basic Datatypes - Numbers, Strings, Boolean
Aggregate Datatypes - Arrays, Structs etc
Refrence Datatypes - Pointers, maps, slices etc
Interface Datatypes - will be discussed later ;-)

All these will be discussed at the time of implementation
*/

package main //"main" is an identifier here and "Package is a keyword"
import "fmt" // Import is a keyword here that we use to prepare our program for execution
func main() { // shows a valid declaration of a function where name of the function is "main"
	var F = "hehe"
	var F1 = 3
	var F2 = 'i'                      // var keyword can be used to declare a variable that can store any kind of datatype
	fmt.Printf("%s %d %c", F, F1, F2) // formatted print statement can be used while printing for better understanding.
	/*There are total 25 keywords available in Golang.*/
	// We can also use int, float to declare variables but you already know how to do that ;-)
	a := 5
	aa := "Hello"
	aaa := '+'                      // Direct assignment (without declaration is possible)
	fmt.Print(a, " ", aa, " ", aaa) // prints 5 Hello AsciiValOf('+')
	var f int8 = 25                 // Syntax to declare and initialize a variable with a specified datatype
	var b int16 = 16275             // Try int32 and int64 on compiler and you'll know how different they are
	fmt.Print(" ", f, b)
	/*Unlike any other programming lanuguage you know, Golang also provides support for
	COMPLEX NUMBERS , Sounds crazy right.*/
	var comp complex64 = complex(5, 9) // first argument represents real and second argument represents imaginary part
	fmt.Print(" ", comp, " ")
	var A bool = true
	if A {
		fmt.Print("true")
	} else {
		fmt.Print("False")
	}
	// You know how this works ---- HEHE

}

/*
This is a basic introduction of various datatypes and keywords used in golang
As we move on, we will see the implementation of more like them with explanation.

Meanwhile Practice on your compiler about the same
*/
