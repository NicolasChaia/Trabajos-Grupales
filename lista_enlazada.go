package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodo := nodoLista[T]{dato: dato, prox: nil}
	return &nodo
}
func CrearListaEnlazada[T any]() Lista[T] {
	lista := listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
		largo:   0,
	}
	return &lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}
