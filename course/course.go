package course

import "fmt"

//sintaxis de una estructura
type Course struct {
	Name string
	Price float64
	IsFree bool
	UserID []uint
	Classes map[uint]string
}
//method change price
func (c *Course) ChangePrice(price float64) {
	// Se le asigna un nuevo valor a la estructura
	c.Price = price
}
//imprime clases
//constructores: funciones que retornan una estructura
func (c Course)PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2]) // borra el último espacio y la coma
}
