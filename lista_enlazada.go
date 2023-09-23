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

// Devuelve si la lista esta vacia o no con un booleano
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}

// Permite insertar un elemento en la primera poscicion de una lista y en tiempo constante.
func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear[T](dato)
	nodo.prox = lista.primero
	lista.primero = nodo
	lista.largo++

}

// Permite insertar un elemento en la ultima poscicio de la lista.
func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear[T](dato)
	lista.ultimo.prox = nodo
	nodo.prox = nil
	lista.ultimo = nodo
	lista.largo++
}

// Se encarga de borra el primer elemento de la lista enlazada y devolverlo.
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.prox
	lista.largo--
	return dato
}

// Permite mostrar el dato del primer elemento de la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

// Permite mostrar el dato del ultimo elemento de la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

// Esta primitiva permite al usuario saber el largo de la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterador Lista

// Permite aplicar la funcion a todos los elementos de la lista, teniendo dos condiciones de corte, que no haya mas elementos en la lista o que se devuelva falso.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		//	visitar(iter)
	}
} // NO SE COMO HACER ESTO

type iterListaEnlazada[T any] struct {
	// Necesitamos el actual, anterior para cuando insertemos o borremos no perdamos las referencias
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// Devuelve un iterador propio de la lista.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := iterListaEnlazada[T]{
		iterador.lista:    lista,
		iterador.actual:   lista.primero,
		iterador.anterior: lista.primero,
	}
	return &iterador
}
