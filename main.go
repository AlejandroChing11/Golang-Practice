package main //Central point - package "#principal - data name"
import "fmt"

//it can be with brackets -- fmt is a package

// func normalFunction(message string) {
// 	fmt.Println(message)
// }

// func tripleArgument(a, b int, c string) {
// 	fmt.Println(a, b, c)
// }

// func returnValue(y int) int {
// 	return y * 2
// }

// func doubleReturnValue(a int) (c, d int) {
// 	return a, a * 2
// }

// func areaRectangulo(base, altura int) int {
// 	return base * altura
// }

func main() {

	//For cycles in Golang.
	For conditional (Tradiocional one)
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// fmt.Printf("\n")

	//For While. Basically it's While because you have to initialize a variable and then start the cycle with the condition and then you have to put the task and the incremental or the function
	// counter := 0
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	//For forever. It's a loop that is not going to stop.
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

	//Reto
	counter := 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}
	// 	normalFunction("Hello world")

	// 	tripleArgument(12, 20, "Alejandro")

	// 	value := returnValue(5)
	// 	fmt.Println(value)

	// 	value1, value2 := doubleReturnValue(2)
	// 	fmt.Println("value1 and value2: ", value1, value2)

	// 	area := areaRectangulo(12, 40)
	// 	fmt.Println("El area del rectangulo es igual a: ", area)

	//How to declare a const in Golang, They are not going to change in the future.
	// const pi float64 = 3.14
	// const pi2 = 3.15
	// fmt.Println("pi1", pi)
	// fmt.Println("pi2", pi2)

	//How to declare int varibles in Golang.
	// base := 12          //This ":" this means the variable isn't declare before.
	// var altura int = 14 //We got to type the data type explicitly
	// var ancho int       //The value of a variable that is initialized but haven't been declared yet is 0.
	// fmt.Println(base, altura, ancho)
	//zero values
	// var a int     //Integer value
	// var b string  //String value
	// var d bool    //True-False boolean
	// var c float64 //Float64 value
	// fmt.Println(a, b, c, d)

	//Area de un cuadrado
	// const baseCuadrado = 12 //camelCase
	// const alturaCuadrado = 14
	// area := baseCuadrado * alturaCuadrado //We use ":=" because we used other variables that were already declared
	// fmt.Println("La base del cuadrado es: ", baseCuadrado, "La altura del cuadrado es: ", alturaCuadrado, "Y el area del cuadrado es: ", area)

	//Let's learn arithmetic operations
	// x := 100
	// y := 10
	// result := x + y
	// resta := x - y
	// division := x / y
	// mod := x % y
	// multiplication := x * y
	//Incremental
	// x++

	// fmt.Println("Suma", result)
	// fmt.Println("Resta", resta)
	// fmt.Println("Division", division)
	// fmt.Println("Modulo", mod)
	// fmt.Println("Multiplication", multiplication)
	// fmt.Println("Incremental", x)

	//decremental
	// x--
	// fmt.Println("Declemental", x)

	//primeral data types
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	//FMT is a package
	//We can interact with the console with this package
	//Println
	// var helloMessage string = "Hello"
	// var worldMesage string = "World"
	// fmt.Println(helloMessage, worldMesage)
	//PrintF
	// nombre := "Alejandro"
	// curso := 500
	//"% is used to data parsing"
	// fmt.Printf("%s tiene mas de %d amigos\n", nombre, curso)
	// fmt.Printf("%v tiene mas de %v amigos\n", nombre, curso)
	//Sprintf
	// message := fmt.Sprintf("%s tiene mas de %d amigos", nombre, curso)
	// fmt.Println(message)
	// Data Types.
	// fmt.Printf("helloMessage: %T \n", helloMessage)
	// fmt.Printf("curso: %T \n", curso)
	// fmt.Printf("nombre: %T \n", nombre)

}
