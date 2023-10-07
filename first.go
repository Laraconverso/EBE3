package main

import (
	"fmt"
)


func main(){

	// imprimit hola mundo 
	fmt.Println("Hola Mundo! :)")

	// declarar variables
	// var nombre string
	// var horas int
	// //asignar valores
	// horas = 20
	// nombre = "Lara"

	// producto1, precio2 := "Jean", 10.5

	// //se puede agrupar en bloques 
	// var (
	// 	producto = "Course" 
	// 	cantidad = 25
	// 	precio = 100
	// 	enStock = true
	// )

	//constantes
	const status = "ok"

	//tipos de datos
	// las variables siempre tienen un tipo de dato definido
	//pueden ser Integers, Floats, Strings, Boolean, Bytes. 
	//uint (sin signo) int(con signo) 

	//arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "world!"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	//slice 
	b := make([]int, 5)
	fmt.Print(b)
	fmt.Print("\n")

	//key-value map
	myMap := map[string]int{}
	myMap2 :=make(map[string]string)
	fmt.Print(myMap)
	fmt.Print("\n")
	fmt.Print(len(myMap2)) //longitud de un map 
	fmt.Print("\n")
	var alumnos = map[string]int{"Benjamin" :23, "Lara": 22}
	fmt.Println(alumnos["Lara"])

}