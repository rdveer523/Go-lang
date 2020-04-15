// DATA TYPES in Go lang
//-------------------------
/*
In Go language we should specify what data type it is and instead ot mentioning the data type
you can do as below by taking optional concise variable declaration and initialization through type inference :=
ex:
x := 0 instead of int x=0, var x=0;
You will find more examples in next files
*/
package main

//package main

import "fmt"

func main() {

	meVeera := true
	YouWho := false
	fmt.Printf("%T\n", meVeera)
	fmt.Printf("%T", YouWho)
}
