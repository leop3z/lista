package listaEnlazada

type Lista[T any] interface {
	// EstaVacia devuelve true si la lista está vacía, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un elemento al comienzo de la lista.
	InsertarPrimero(elem T)

	// InsertarUltimo agrega un elemento al final de la lista. El elemento debe ser valido.
	InsertarUltimo(elem T)

	// BorrarPrimero elimina y devuelve el primer elemento de la lista.
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista sin removerlo.
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista sin removerlo.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recorre la lista desde el primero hasta el último elemento,
	// aplicando la función visitar a cada uno hasta que esta retorne false o
	// finalice la lista.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo para recorrer la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual devuelve el elemento en la posición actual del iterador.
	VerActual() T

	// HaySiguiente indica si hay un elemento en la posición actual del iterador.
	HaySiguiente() bool

	// Siguiente avanza el iterador a la siguiente posición.
	Siguiente()

	// Insertar agrega un elemento en la posición actual del iterador.
	Insertar(elem T)

	// Borrar elimina y devuelve el elemento en la posición actual del iterador.
	Borrar() T
}
