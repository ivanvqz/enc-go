package main

import (
	"fmt"

	"github.com/ivanvqz/enc-go/course"

)


func main() {
	// Función que inicializa una estructura
	Go := course.NewCourse("Curso de flask", 0, false)

	// Se le asigna un nuevo valor a la estructura
	Go.SetUserID( []uint{1, 2, 3})
	fmt.Println(Go.UserID())
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})
	Go.SetName("")
	fmt.Println(Go.Name())

	// Go.PrintClasses()
	// React := Course{Name: "Curso de React",}
	// fmt.Println(Go.Name)
	// fmt.Printf("%+v", React)
	//se crea un puntero a la estructura
	// Go.ChangePrice(20.00)
	// fmt.Println(Go.Price)
	// Go.Classes[4] = "Slices"
	Go.PrintClasses()
	// fmt.Printf("%+v", Go)	
}
