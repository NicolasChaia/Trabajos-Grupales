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
	if lista.primero == nil {
		//Caso en donde la lista esta vacia
		lista.primero = nodo
		lista.primero.prox = lista.ultimo
		lista.ultimo = nodo
	}
	nodo.prox = lista.primero
	lista.primero = nodo
	lista.largo++

}

// Permite insertar un elemento en la ultima poscicion de la lista.
func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear[T](dato)
	if lista.primero == nil {
		//Caso en donde la lista esta vacia
		lista.primero = nodo
		lista.primero.prox = lista.ultimo
		lista.ultimo = nodo
	} else {
		lista.ultimo.prox = nodo
		lista.ultimo = nodo
	}
	lista.largo++
}

// Se encarga de borrar el primer elemento de la lista enlazada y devolverlo.
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verSiEstaVacia()
	dato := lista.primero.dato
	if lista.primero == lista.ultimo {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		lista.primero = lista.primero.prox
	}
	lista.largo--
	return dato
}

// Permite mostrar el dato del primer elemento de la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verSiEstaVacia()
	return lista.primero.dato
}

// Permite mostrar el dato del ultimo elemento de la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verSiEstaVacia()
	return lista.ultimo.dato
}

// Esta primitiva permite al usuario saber el largo de la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) verSiEstaVacia() {
	// Esta funcion verifica si la lista esta vacia, si lo esta, entra en panico
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

// Iterador Lista

type iterListaEnlazada[T any] struct {
	// Necesitamos el actual, anterior para cuando insertemos o borremos no perdamos las referencias
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// Devuelve un iterador propio de la lista.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := iterListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
	return &iter
}

// Permite aplicar la funcion a todos los elementos de la lista, teniendo dos condiciones de corte, que no haya mas elementos en la lista o que se devuelva falso.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		valor := iter.VerActual()
		if !visitar(valor) {
			break
		}
		iter.Siguiente()
	}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	} else {
		return iter.actual.dato
	}
}
func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}
func (iter *iterListaEnlazada[T]) Siguiente() {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	} else {
		iter.anterior = iter.actual
		iter.actual = iter.actual.prox
	}
}
func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	elem := nodoCrear[T](elemento)
	if iter.lista.EstaVacia() {
		//Caso en que esta vacia
		iter.lista.primero = elem
		iter.lista.ultimo = elem
	} else if iter.anterior == nil {
		//Caso donde se inserta primero
		elem.prox = iter.actual
		iter.lista.primero = elem
		iter.actual = elem

	} else if iter.actual == nil {
		//Caso donde se quiera insertar al ultimo
		iter.anterior.prox = elem
		elem.prox = nil
		iter.actual = elem
		iter.lista.ultimo = elem
	} else {
		iter.anterior.prox = elem
		iter.lista.ultimo = elem
	}
	iter.lista.largo++
	iter.actual = elem
}
func (iter *iterListaEnlazada[T]) Borrar() T {
	// Terminar de chequear esta funcion integra
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	elemento := iter.actual.dato

	if iter.anterior == nil {
		//Caso de que se elimina el primero
		iter.lista.primero = iter.actual.prox
		iter.actual = iter.actual.prox
	} else if iter.actual.prox == nil {
		//Caso en que se quiera eliminar el ultimo
		iter.anterior.prox = iter.actual.prox
		iter.lista.ultimo = iter.anterior
		iter.actual = nil
	} else if iter.lista.primero == iter.actual && iter.lista.ultimo == iter.actual {
		//Caso en donde queda un solo elemento
		iter.actual = nil
		iter.lista.primero = iter.actual
		iter.lista.ultimo = iter.actual

	} else {
		// Caso general, elementos en el "medio" de la lista
		iter.actual = iter.actual.prox
		iter.anterior.prox = iter.actual
	}
	iter.lista.largo--
	return elemento
}
