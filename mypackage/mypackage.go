package mypackage

// Interfaz para metodos que comparten el mismo nombre
type ADT interface {
	EsVacio() bool
}

//ESTRCUTURA DEL NODO Aplicable a la Lista, Pilas y Colas simples

type Nodo struct {
	valor     int
	siguiente *Nodo
}

// SetValor establece el nuevo valor del nodo
func (n *Nodo) SetValor(dato int) {
	n.valor = dato
}

// GetValor devuelve el valor del nodo
func (n *Nodo) GetValor() int {
	return n.valor
}

// SetSiguiente establece el siguiente nodo en la lista
func (n *Nodo) SetSiguiente(elSigue *Nodo) {
	n.siguiente = elSigue
}

// GetSiguiente devuelve el siguiente nodo en la lista
func (n *Nodo) GetSiguiente() *Nodo {
	return n.siguiente
}
