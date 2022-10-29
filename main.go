package main
import "fmt"

func main() {

	Go := Course{
		Name: "Curso de Go",
		Price: 10.00,
		IsFree: false,
		UserID: []uint{1,2,3},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}
	// Go.PrintClasses()
	// React := Course{Name: "Curso de React",}
	// fmt.Println(Go.Name)
	// fmt.Printf("%+v", React)
	//se crea un puntero a la estructura
	Go.ChangePrice(20.00)
	fmt.Println(Go.Price)
}