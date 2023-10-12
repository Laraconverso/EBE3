package c3

import "fmt"

func main(){
	// Estructuras:
	type Persona struct{
		Nombre string
		Apellido string
		Edad int
		Profesion string
	}

	//instancio una estructura
	p1:= Persona{"Chapulin", "Colorado", 1, "Animador"}
	fmt.Print(p1)

	p2:= Persona{
		Nombre:  "nombre",
		Apellido: "Apellido",
		Edad: 14,
		Profesion: "Medico",
	}
	fmt.Print(p2)
}