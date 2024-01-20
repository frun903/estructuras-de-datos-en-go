package mypackage

import (
	"fmt"
)

// ESTRUCTURA DE LA LISTA
type Lista struct {
	Inicio *Nodo
}

func (l *Lista) EsVacio() bool {
	if l.Inicio == nil {
		return true
	} else {
		return false
	}

}

// Insertar inserta un nuevo nodo con el valor dado en la posición indicada de la lista.
// Si la posición es 0, el nuevo nodo se agrega al principio de la lista.
//
// Parámetros:
//   - pos: posición en la que insertar el nuevo nodo.
//   - valor: valor del nuevo nodo a insertar.
func (l *Lista) Insertar(pos, valor int) {
	posActual := 0

	nodoAux := l.Inicio
	nuevo := &Nodo{} // Creo un nuevo nodo con el {} seria el equivalente al new de c++

	nuevo.SetValor(valor)

	if pos == 0 {
		nuevo.SetValor(valor)
		nuevo.SetSiguiente(l.Inicio)
		l.Inicio = nuevo
		return
	}

	for {

		if nodoAux.GetSiguiente() == nil {
			//si entra aqui es porque llegamos a que el siguiente es null
			nodoAux.SetSiguiente(nuevo)
			break
		}

		if posActual+1 != pos {
			nodoAux = nodoAux.GetSiguiente()
			posActual++
		} else {
			break
		}

	}

	nuevo.SetSiguiente(nodoAux.GetSiguiente())
	nodoAux.SetSiguiente(nuevo)

}

func (l *Lista) ObtenerValor(pos int) int {
	nodoAux := l.Inicio
	//posActual := 0

	for i := 0; i < pos; i++ {
		if nodoAux.GetSiguiente() == nil {
			break

		}
		nodoAux = nodoAux.GetSiguiente()
	}

	return nodoAux.GetValor()
}

// Obtener el tamanio una lista
//
// La lista tiene un tamanio 1->2->3
// El tamanio es 3
//
// Sin paremetros de entrada, de salida devuelve un entero
func (l *Lista) TamanioLista() int {

	//Vamos a hacer una verificacion temprana para ver que
	//mi lista no este vacia esto se hara antes que nada para
	// no generar errores
	if l.Inicio == nil {
		// Si la lista está vacía, devuelve 0
		return 0
	}

	aux := *l.Inicio //puntero a nodo

	tamanio := 0

	//Voy a usar un For Forever y voy a remperlo
	//Con una condicion cuando el proximo sea nullptr o nill
	for {

		if aux.GetSiguiente() == nil {
			//El proximo ya es nill por lo tanto sumo uno mas y rompo el ciclo
			tamanio++
			break
		}
		tamanio++
		aux = *aux.GetSiguiente()

	}
	return tamanio

}

// Imprimir toma una lista
//
// La lista es impresa de la forma 1->2->3->nullptr
//
// Sin paremetros de entrada ni de salida
func (l *Lista) Imprimir() {
	//Vamos a hacer una verificacion temprana para ver que
	//mi lista no este vacia esto se hara antes que nada para
	// no generar errores
	if l.Inicio == nil {
		// Si la lista está vacía, devuelve 0
		fmt.Println("La lista esta vacia")
	} else {

		for i := 0; i < l.TamanioLista(); i++ {
			fmt.Print(l.ObtenerValor(i), "->")
		}
		fmt.Print("nullptr\n")
	}
}

// Toma una lista y elimina su nodo
//
// La lista va del 0 a n, debo pasar una una pocision desde 0
// a n para eleminar el nodo en esa pocision
//
// Parametro de entrada entero que indica pocision de la lista, sin parametros de salida
func (l *Lista) DeleteNodo(pos int) {
	if pos >= l.TamanioLista() || pos < 0 {
		fmt.Println("Posición ingresada no válida")
		return
	}

	//Si debo eliminar la primera pocision
	if pos == 0 {

		l.Inicio = l.Inicio.GetSiguiente()
		return
	}

	posActual := 0
	//Creo el nodo auxiliar para recorrer la lista
	nodoAux := l.Inicio
	//Voy a crea un nodo auxiliar mas
	nodoAuxQuevoyaBorrar := l.Inicio

	for {

		if posActual+1 != pos {
			nodoAux = nodoAux.GetSiguiente()
			posActual++
		} else {
			break
		}

	}

	//digo que el a borrar va a ser igual al siguiente donde se pociono Aux
	nodoAuxQuevoyaBorrar = nodoAux.GetSiguiente()

	// Si el nodo a borrar es el último
	if nodoAuxQuevoyaBorrar.GetSiguiente() == nil {
		nodoAux.SetSiguiente(nil)

	}

	//Para borrarlo solo lo hago que su anterior apunte al proximo y se pierda en memoria
	nodoAux.SetSiguiente(nodoAuxQuevoyaBorrar.GetSiguiente())

}
