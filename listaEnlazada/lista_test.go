package listaEnlazada_test

import (
	"github.com/stretchr/testify/require"
	"lista/listaEnlazada"
	"testing"
)

func TestListaVacia(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

func TestInsertarPrimero(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	lista.InsertarPrimero(10)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
}

func TestInsertarUltimo(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	valor := lista.BorrarPrimero()
	require.Equal(t, "a", valor)
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "b", lista.VerPrimero())
	valor = lista.BorrarPrimero()
	require.Equal(t, "b", valor)
	require.True(t, lista.EstaVacia())
}

func TestIterar(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	suma := 0
	lista.Iterar(func(elem int) bool {
		suma += elem
		return true
	})
	require.Equal(t, 6, suma)
}

func TestCorteEnIteradorInterno(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i)
	}
	suma := 0
	lista.Iterar(func(val int) bool {
		suma += val
		return val != 2
	})
	require.Equal(t, 3, suma)
}

func TestIteradorInsertarYVerActual(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(100)
	require.Equal(t, 1, lista.Largo())
	iter2 := lista.Iterador()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 100, iter2.VerActual())
}

func TestIteradorSiguiente(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestInsertarAlInicioConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(42)
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 42, lista.VerPrimero())
}

func TestInsertarAlFinalConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(3)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestInsertarEnElMedioConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)
	expect := []int{1, 2, 3}
	i := 0
	lista.Iterar(func(val int) bool {
		require.Equal(t, expect[i], val)
		i++
		return true
	})
}

func TestBorrarPrimeroConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	iter := lista.Iterador()
	valor := iter.Borrar()
	require.Equal(t, "a", valor)
	require.Equal(t, "b", lista.VerPrimero())
}

func TestBorrarUltimoConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[string]()
	lista.InsertarUltimo("x")
	lista.InsertarUltimo("y")
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar("z")
	iter2 := lista.Iterador()
	for iter2.HaySiguiente() {
		iter2.Siguiente()
	}
	iter2.Borrar()
	require.Equal(t, "y", lista.VerUltimo())
}

func TestBorrarEnElMedioConIterador(t *testing.T) {
	lista := listaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	valor := iter.Borrar()
	require.Equal(t, 2, valor)
	expect := []int{1, 3}
	i := 0
	lista.Iterar(func(val int) bool {
		require.Equal(t, expect[i], val)
		i++
		return true
	})
}
