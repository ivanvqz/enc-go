package course

import "fmt"

//sintaxis de una estructura
type course struct {
	name string
	price float64
	isFree bool
	userID []uint
	classes map[uint]string
}
//setter name
func (c *course) SetName(name string) {
	// Se le asigna un nuevo valor a la estructura si es que no se especifica
	if name == "" {
		c.name = "Sin nombre"
		return
	}
	// Se le asigna un nuevo valor a la estructura
	c.name = name
}

//getter name
func (c course) Name() string {
	return c.name
}

//getter price
func (c *course) SetPrice(price float64) { c.price = price }
func (c course) Price() float64 { return c.price  }

// setter isfree
func (c *course) SetIsFree(isFree bool) { c.isFree = isFree }
func (c course) IsFree() bool { return c.isFree }

//seter userid
func (c *course) SetUserID(userID []uint) { c.userID = userID }
func (c course) UserID() []uint { return c.userID }

//setter
func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}


func NewCourse(name string, price float64, isfree bool) *course {
	if price == 0 {
		price = 30.00
	}

	return &course{
		//se asigna el valor de la variable name a la estructura
		name: name,
		price: price,
		isFree: isfree,
	}
}
//imprime clases
//constructores: funciones que retornan una estructura
func (c course) PrintClasses() {
	text := "Las clases son: "

	for _, class := range c.classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2]) // borra el último espacio y la coma
}
