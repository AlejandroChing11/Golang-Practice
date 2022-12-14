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

// func par(num int) {
// 	if num%2 == 0 {
// 		fmt.Println("Es par")
// 	} else {
// 		fmt.Println("Es impar")
// 	}
// }

// func password(user, passwords string) {
// 	if user != "" && passwords != "" {
// 		fmt.Println("El usuario ", user, "Es valido y la contraseña ", passwords, "Es valida")
// 	} else {
// 		fmt.Println("Inteneta de nuevo")
// 	}
// }

// func isPalindrome(text string) {
// 	var textReverse string
// 	for i := len(text) - 1; i >= 0; i-- {
// 		textReverse += string(text[i])
// 	}

// 	if text == textReverse {
// 		fmt.Println("Es palindromo")
// 	} else {
// 		fmt.Println("No es palindromo")
// 	}

// }

func main() {
	// isPalindrome(strings.ToLower("Reconocer"))

	//For cycles in Golang.
	// For conditional (Tradiocional one)
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Printf("\n")

	//For While. Basically it's While because you have to initialize a variable and then start the cycle with the condition and then you have to put the task and the incremental or the function
	// counter := 0
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	//For forever. It's a loop that is not going to stop.
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	//Reto
	// counter := 10
	// for counter > 0 {
	// 	fmt.Println(counter)
	// 	counter--
	// }
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

	//primeral data types\clases\2232-programacion-golang\35750-llave-valor-con-maps\
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

	//Conditionals in Golang.
	// valor1 := 1
	// valor2 := 1
	// if valor1 == 1 {
	// 	fmt.Println("Es igual a 1")
	// } else {
	// 	fmt.Println("No es igual a 1")
	// }

	// With &&
	// if valor1 == 1 && valor2 == 1 {
	// 	fmt.Println("Es verdad")
	// }

	// With ||
	// if valor1 == 1 || valor2 == 2 {
	// 	fmt.Println("Es verdad")
	// }

	//Convert string to int
	//nill indica si una funcion indicada no tiene un error.
	//value = va a guardar el valor y error = va a guardar el error siempre y cuandoe exista
	// value, err := strconv.Atoi("100")
	// if err != nil { //nill means when a function has an error
	// log.Fatal(err) //This is to print the error message
	// }
	// fmt.Println(value) //This is to print the value converted to int

	//Reto
	// par(2)
	// password("Alejandro", "Alejandro1102")

	//Learning switch

	// switch modulo := 4 % 2; modulo { //We use this switch when we are going to evaluate over a unique variable.
	// case 0:
	// 	fmt.Println("Es par")
	// default:
	// 	fmt.Println("Es impar")
	// }

	//Switch Without conditions
	// value := 50
	// switch {
	// case value > 100: //We use this switch when we are going to add multiple conditions
	// 	fmt.Println("Es mayor que 100")
	// case value < 0:
	// 	fmt.Println("Es menor a 0")
	// default:
	// 	fmt.Println("No condition")
	// }

	//The use of defer, break and continue
	//Defer.
	// defer fmt.Println("Hola") //We use defer to keep the function to the final, so we can use it when we are connecting to a database and we want to shut douwn the connection
	// fmt.Println("Mundo")

	//Continue break (We usse both on for)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)

	//continue
	// if i == 2 {
	// fmt.Println("Es dos")
	// continue //We use continue when a the condition on the for contains something that is interesting for us and we want to continue the loop. For example when we have an error on the loop but it is controlled so we can move on the code.
	// }

	//Break
	// if i == 8 {
	// fmt.Println("Break")
	// break //We use break to broke the loop.
	// }

	// }

	//Learning Arrays in Golang - Go (QThey are unmutable)
	// var array [4]int
	// array[0] = 1
	// array[1] = 2
	// fmt.Println(array, len(array), cap(array))
	// fmt.Printf("\n")
	// var nombres [3]string
	// nombres[0] = "Jenzey"
	// nombres[1] = "Alejandro"
	// nombres[2] = "Mei li"
	// fmt.Println(nombres, len(nombres), cap(nombres)) //len = longitud del array , cap = array capacity (those are important functions)
	// fmt.Printf("\n")

	//Slice  (They can be mutable)
	// slice := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(slice, len(slice), cap(slice))
	// fmt.Printf("\n")

	//Slicing
	//Methods
	// fmt.Println(slice[0])
	// fmt.Println(slice[:3])
	// fmt.Println(slice[2:4])
	// fmt.Println(slice[4:])
	// fmt.Printf("\n")

	//Append
	// slice = append(slice, 6)
	// fmt.Println(slice)
	// fmt.Printf("\n")

	// Apend with a new list.
	// newSlice := []int{7, 8, 9, 10}
	// slice = append(slice, newSlice...) //"..." means that Go is going to add the new elements independently
	// fmt.Println(slice)

	// slice := []string{"Hola", "que", "haces"} //This is not an array, this is a slice.

	// for _, value := range slice { //We use the for loop. the first thing you have to initialize a variable and then a value that is the one that is going to have the values of the slice
	// fmt.Println(value)
	// }

	//Maps in Golang such as dictionaries in python and objects in javascript
	m := make(map[string]int) //String is the value, int is the key
	m["Jose"] = 14
	m["Alejandro"] = 20
	m["Maria"] = 14

	fmt.Println(m)

	//Recorrer un map with range
	for i, value := range m {
		fmt.Println(i, value)
	}

	fmt.Printf("\n")

	//Find a value.
	value, sure := m["Jose"] //var "ok" is a variable that indicates if the value exists in the map (it can has whatever name)
	fmt.Println(value, sure)

}
