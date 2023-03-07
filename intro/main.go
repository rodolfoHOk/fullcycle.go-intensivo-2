package main

import "fmt"

func main_old_1() {
	print("Hello, world!") // imprime Hello, world! no terminal
}

func main_old_2() {
	a := "Rudolf" // ":=" declara e atribui valor a variável a
	print(a, "\n")
	a = "X" // atribui valor a variável a
	print(a, "\n")

	var b string // declara variável b
	b = "Y" // atribui valor a variável b
	print(b, "\n")

	// inferência de tipo e tipagem estática
	c := 1
	d := 1.0
	e := true
	f := false
	print(c, "\n")
	print(d, "\n")
	print(e, "\n")
	print(f, "\n")
}

// if / else
func main_old_3() {
	a := 10
	b := 11

	if a > b {
		print("a > b", "\n")
	} else if a < b {
		print("a < b", "\n")
	} else {
		print("a == b", "\n")
	}
}

// obs: go só tem for não tem while, do while, etc...

func main_old_4() {
	fmt.Println("Hello, world!") // fmt é um pacote do go e deve ser importado antes
}

// Struct
type Order struct {
	ID string
	Price float64
	Quantity int
}

// Método da struct Order
func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

// Método da struct Order sem apontamento
func (o Order) setPriceWithoutPointer(price float64) {
	o.Price = price
	fmt.Println("Price dentro do set price:", o.Price)
}

// Método da struct Order com apontamento "*Order"
func (o *Order) setPriceWithPointer(price float64) {
	o.Price = price
	fmt.Println("Price dentro do set price com apontamento:", o.Price)
}

func main_old_5() {
	order := Order{
		ID: "123",
		Price: 10.0,
		Quantity: 5,
	}

	fmt.Println(order.ID, order.Price, order.Quantity)

	order.setPriceWithoutPointer(20.0)

	fmt.Println("Preço total:", order.getTotal())

	fmt.Println("Price:", order.Price) // o set price não mudou o valor de order price

	order.setPriceWithPointer(20.0)

	fmt.Println("Preço total:", order.getTotal())
}

// ponteiros
func main_old_6() {
	a := 10
	b := a
	fmt.Println(b) // 10
	b = 20
	fmt.Println(a) // 10
	fmt.Println(b) // 20 => a e b são independentes

	c := &a // referência de memória
	fmt.Println(c) // 0xc0000b4000 = c é um endereço de memória que correspondente ao endereço de a
	fmt.Println(*c) // 10 = valor no endereço de memória c

	a = 20
	fmt.Println(*c) // 20

	// obs: sempre que ver um "*" quer dizer que está apontando ao valor do endereço
	// ao invés de copiar o valor, e sempre que mudar a variável com "*" vai mudar todas as
	// variáveis que tenham o mesmo apontamento
}

func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do set price com apontamento:", o.Price)
}

// Construtor sem ponteiro
// func NewOrder() Order {
// 	return Order{}
// }

// construtor com ponteiro => muito comum no Go economiza memória e evita duplicidade
func NewOrder() *Order {
	return &Order{}
}

func main() {
	order := NewOrder()
	order.ID = "123"
	order.Price = 10.0
	order.Quantity = 5

	order2 := order; // com construtor sem ponteiro cria uma cópia
	order.Price = 5
	fmt.Println("Preço order 2:", order2.Price) // com construtor com ponteiro muda o valor

	fmt.Println(order, order2)
}
