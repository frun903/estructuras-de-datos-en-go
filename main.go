package main

import (
	"fmt"

	adt "mypackage/mypackage"
)

func main() {

	fmt.Println("Hello! Time of Strcut of Data!")

	fmt.Println("\nEjemplo de Uso de una Lista enlazada en Go")

	var unaLista adt.Lista

	unaLista.Insertar(0, 15)
	unaLista.Insertar(0, 26)
	unaLista.Insertar(0, 86)

	fmt.Println(unaLista.TamanioLista())
	unaLista.Imprimir()

	fmt.Println("\nColoco un nuevo valor en medio de la lista")

	unaLista.Insertar(2, 105)
	unaLista.Imprimir()

	unaLista.DeleteNodo(2)
	unaLista.Imprimir()

	fmt.Println("\n\nEjemplo de Uso de una Pila en Go")

	var unaPila adt.Pila

	if unaPila.EsVacio() == true {
		fmt.Println("La Pila es Vacia")
	} else {
		fmt.Println("La pila no es vacia")
	}

	unaPila.Push(15)
	unaPila.Push(26)
	unaPila.Push(89)

	if unaPila.EsVacio() == true {
		fmt.Println("La Pila es Vacia")
	} else {
		fmt.Println("La pila no es vacia")
	}

	fmt.Println(unaPila.Pop())
	fmt.Println(unaPila.Pop())
	fmt.Println(unaPila.Pop())
	fmt.Println(unaPila.Pop())

	fmt.Println("\n\nEjemplo de Uso de una Cola en Go")
	var unaQueue adt.Cola

	unaQueue.Encolar(85)
	unaQueue.Encolar(505)
	unaQueue.Encolar(72)

	if unaQueue.EsVacio() == true {
		fmt.Println("La Queue esta Vacia")
	} else {
		fmt.Println("La queue no es vacia")
	}

	fmt.Println(unaQueue.Desencolar())
	fmt.Println(unaQueue.Desencolar())

}
