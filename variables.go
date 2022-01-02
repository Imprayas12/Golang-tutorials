// This program will implement the use of various identifiers,datatypes,variables and keywords which
//are available in Go programming language so let;s GOPHER it.

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
}
