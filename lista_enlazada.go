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
	return lista.primero == nil && lista.ultimo == nil && lista.Largo() == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear[T](dato)
	if lista.EstaVacia() {
		// Caso en donde la lista esta vacia
		lista.primero = nodo
		lista.primero.prox = lista.ultimo
		lista.ultimo = nodo
	} else {
		nodo.prox = lista.primero
	}
	lista.primero = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear[T](dato)
	if lista.EstaVacia() {
		// Caso en donde la lista esta vacia
		lista.primero = nodo
		lista.primero.prox = lista.ultimo
		lista.ultimo = nodo
	} else {
		lista.ultimo.prox = nodo
		lista.ultimo = nodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verSiEstaVacia()
	dato := lista.primero.dato
	// Caso en el que la lista tiene un solo elemento
	if lista.primero == lista.ultimo {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		lista.primero = lista.primero.prox
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verSiEstaVacia()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verSiEstaVacia()
	return lista.ultimo.dato
}

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
	actual := lista.primero
	for actual != nil {
		valor := actual.dato
		if !visitar(valor) {
			break
		}
		actual = actual.prox
	}
}

// Esta funcion devuelve el dato en el que esta posicionado el iterador en esa la iteracion actual
func (iter *iterListaEnlazada[T]) VerActual() T {
	if iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	} else {
		return iter.actual.dato
	}
}

// Esta funcion devuelve un booleano si hay mas elementos para seguir iterando en la lista
func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

// Esta funcion, siempre y cuando exista, "avanza" al siguiente nodo
func (iter *iterListaEnlazada[T]) Siguiente() {
	if iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	} else {
		iter.anterior = iter.actual
		iter.actual = iter.actual.prox
	}
}

// Esta funcion se encarga de insertar un elemento en la posicion acutual del iterador

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	elem := nodoCrear[T](elemento)
	if iter.anterior == nil {
		if iter.actual != nil {
			elem.prox = iter.actual
		} else {
			iter.lista.ultimo = elem
		}
		iter.lista.primero = elem
	}
	if iter.anterior != nil {
		if iter.actual == nil {
			iter.lista.ultimo = elem
		} else {
			elem.prox = iter.actual
		}
		iter.anterior.prox = elem
	}
	iter.actual = elem
	iter.lista.largo++
}

/*Hay errores importantes en el codigo:
Código repetido: Se debe reutilizar EstaVacia en InsertarPrimero
Código repetido: Se debe reutilizar EstaVacia en InsertarUltimo
Advertencia (correccion no obligatoria, pero importante): Se debería reutilizar HaySiguiente en VerActual
Advertencia (correccion no obligatoria, pero importante): Se debería reutilizar HaySiguiente en Siguiente
Error: El iteardor interno no debe depender del iterador externo! */
// Esta funcion se encarga de borrar un elemento en la posicion acutual del iterador
func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.VerActual()
	// Caso en que quede un solo elemento
	if iter.lista.largo == 0 {
		iter.lista.primero = nil
		iter.lista.ultimo = nil
	}
	// Caso de que se quiera borrar el primero de la lista
	if iter.anterior == nil && iter.lista.largo != 0 {
		iter.lista.primero = iter.actual.prox

	} else {
		iter.anterior.prox = iter.actual.prox
	}
	// Caso de que se quiera borra el ultimo de la lista
	if iter.actual.prox == nil && iter.lista.largo != 0 {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.prox
	iter.lista.largo--
	return elemento
}
