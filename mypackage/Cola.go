package mypackage

import "fmt"

//COLA El Tipos de Dato Abstracto-Abstract data Type-ADT  COLA es una colección ordenada de elementos de la que se pueden borrar elementos en un extremo (FRENTE de la cola)
// o insertarlos en el otro (FINAL de la cola). Este _abstract data type_
// utiliza la propiedad **FIFO first in first aut primero entrar primero en salir**
//Siempre sale primero el que entro.
//Sus funciones son similares a las de una pila pero el
//funcionamiento es distinto.

// ESTRUCTURA DE una Pila
type Cola struct {
	Tope  *Nodo //inicio de la cola
	Fondo *Nodo //final
}

// EsVacio verifica si la cola está vacía.
// Retorna true si está vacía, de lo contrario, retorna false.
func (c *Cola) EsVacio() bool {
	if c.Tope == nil {
		return true
	} else {
		return false
	}

}

// Encolar agrega un nuevo nodo con el valor proporcionado al final de la cola.
//
// Parámetro:
//   - data: valor del nuevo nodo a encolar.
func (c *Cola) Encolar(data int) {
	nuevoNodo := &Nodo{}
	nuevoNodo.SetValor(data)

	if c.EsVacio() {
		//En el caso de ser vacio que el nuevo nodo sea el tope
		c.Tope = nuevoNodo
	} else {
		//Sino que guarde el nuevo nodo despues del fondo
		c.Fondo.SetSiguiente(nuevoNodo)
	}

	//y que nuevo nodo sea el Fonde acutal
	c.Fondo = nuevoNodo

}

// Desencolar elimina y devuelve el valor del nodo en el tope de la cola.
// Retorna 0 si la cola está vacía.
func (c *Cola) Desencolar() int {
	if c.EsVacio() {
		fmt.Println("Cola ingresada es Vacia no es posible desencolar")
		return 0
	}

	varAux := c.Tope.GetValor()
	c.Tope = c.Tope.GetSiguiente()

	return varAux

}
