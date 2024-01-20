package mypackage

import "fmt"

// PILA Una pila es una colección ordenada de elementos
//en la que pueden _insertarse y suprimirse elementos por
//un extremo llamado TOPE.
//Es como si se apilaran cosas y cuando saco debo sacar la ultima
// en entrar, esto se lo conoce como **LIFO**
//Last in first out significa Ultimo entrar primero en salir.

// ESTRUCTURA DE una Pila
type Pila struct {
	Tope *Nodo
}

// Push agrega un nuevo nodo con el valor proporcionado a la cima de la pila.
//
// Parámetro:
//   - data: valor del nuevo nodo a agregar.
func (p *Pila) Push(data int) {
	nuevoNodo := &Nodo{}
	nuevoNodo.SetValor(data)

	//Que apunte al tope viejo para no eliminar todos los elementos
	//de la pila
	nuevoNodo.SetSiguiente(p.Tope)

	p.Tope = nuevoNodo

}

// Pop elimina y devuelve el valor del nodo en la cima de la pila.
// Retorna 0 si la pila está vacía.
func (p *Pila) Pop() int {
	if p.EsVacio() {
		fmt.Println("Es vacia! no se puede seguir sacando")
		return 0
	}

	auxNodo := p.Tope
	var varAux int

	varAux = p.Tope.GetValor()

	//si recordamos dije que el Aux era igual al tope
	//Ahora lo hago igual al siguiente
	auxNodo = auxNodo.GetSiguiente()

	p.Tope = auxNodo

	return varAux

}

// EsVacio verifica si la pila está vacía.
// Retorna true si está vacía, de lo contrario, retorna false.
func (p *Pila) EsVacio() bool {
	if p.Tope == nil {
		return true
	} else {
		return false
	}
}
