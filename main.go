package main //Central point - package "#principal - data name"

import "fmt" //it can be with brackets -- fmt is a package

func main() {
	//How to declare a const in Golang, They are not going to change in the future.
	const pi float64 = 3.14
	const pi2 = 3.15
	fmt.Println("pi1", pi)
	fmt.Println("pi2", pi2)

	//How to declare int varibles in Golang.
	base := 12          //This ":" this means the variable isn't declare before.
	var altura int = 14 //We got to type the data type explicitly
	var ancho int       //The value of a variable that is initialized but haven't been declared yet is 0.
	fmt.Println(base, altura, ancho)
}
