package listaEnlazada

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{largo: 0}
}

func crearNodo[T any](data T) *nodoLista[T] {
	return &nodoLista[T]{dato: data}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodo_aux := crearNodo(elem)
	if lista.EstaVacia() {
		lista.primero = nodo_aux
		lista.ultimo = nodo_aux
	} else {
		nodo_aux.siguiente = lista.primero
	}
	lista.primero = nodo_aux
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodo_aux := crearNodo(elem)
	if lista.EstaVacia() {
		lista.ultimo = nodo_aux
		lista.primero = nodo_aux
	} else {
		lista.ultimo.siguiente = nodo_aux
	}
	lista.ultimo = nodo_aux
	lista.largo++
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista está vacía")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista está vacía")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista está vacía")
	}
	elem := lista.VerPrimero()

	if lista.largo == 1 {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		lista.primero = lista.primero.siguiente
	}

	lista.largo--
	return elem
}

// ITERADOR INTERNO

func (lista listaEnlazada[T]) Iterar(visitar func(i T) bool) {
	nodo_actual := lista.primero

	for nodo_actual != nil && visitar(nodo_actual.dato) {
		nodo_actual = nodo_actual.siguiente
	}
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

// ITERADOR EXTERNO

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := &iterListaEnlazada[T]{}
	iterador.lista = lista
	iterador.actual = lista.primero

	return iterador
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("Ya no hay nada para iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("Ya no hay nada para iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(t T) {
	nodo := crearNodo(t)
	nodo.siguiente = iterador.actual

	if iterador.anterior == nil {
		iterador.lista.primero = nodo
	}

	iterador.actual = nodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if iterador.actual == nil {
		panic("La lista está vacía")
	}

	dato := iterador.actual.dato

	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}

	iterador.actual = iterador.actual.siguiente
	iterador.lista.largo--

	return dato
}
