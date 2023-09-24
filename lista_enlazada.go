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

// Permite insertar un elemento en la ultima poscicion de la lista.
func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear[T](dato)
	lista.ultimo.prox = nodo
	nodo.prox = nil
	lista.ultimo = nodo
	lista.largo++
}

// Se encarga de borrar el primer elemento de la lista enlazada y devolverlo.
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	// if lista.EstaVacia() {
	// 	panic("La lista esta vacia")
	// }
	lista.verSiEstaVacia()
	dato := lista.primero.dato
	lista.primero = lista.primero.prox
	lista.largo--
	return dato
}

// Permite mostrar el dato del primer elemento de la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	// if lista.EstaVacia() {
	// 	panic("La lista esta vacia")
	// }
	lista.verSiEstaVacia()
	return lista.primero.dato
}

// Permite mostrar el dato del ultimo elemento de la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	// if lista.EstaVacia() {
	// 	panic("La lista esta vacia")
	// }
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

// Permite aplicar la funcion a todos los elementos de la lista, teniendo dos condiciones de corte, que no haya mas elementos en la lista o que se devuelva falso.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	// for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {

	// }
	// La verdad que no me acuerdo muy bien como usar ese ciclo for (me acuerdo que a Martin le gustaba mas pero yo creo que asi podria andar)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		valor := iter.Siguiente()
		// no me acuerdo si habia que cortar cuando sea true o cuando sea false, en este caso lo puse por si es false-> termino la iteracion
		if !visitar(valor) {
			break
		}

	}
}

type iterListaEnlazada[T any] struct {
	// Necesitamos el actual, anterior para cuando insertemos o borremos no perdamos las referencias
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// Devuelve un iterador propio de la lista.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := iterListaEnlazada[T]{
		// iterador.lista:    lista,
		// iterador.actual:   lista.primero,
		// iterador.anterior: lista.primero,
		// ¿No seria asi en realidad?:
		actual:   lista.primero,
		anterior: lista.primero,
		lista:    lista,
	}
	return &iterador
}
func (lista *listaEnlazada[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	} else {
		return iter.actual.dato
	}
}
func (lista *listaEnlazada[T]) HaySiguiente() bool {
	return iter.actual.prox != nil
}
func (lista *listaEnlazada[T]) Siguiente() {
	if iter.actual.prox == nil {
		panic("El iterador termino de iterar")
	} else {
		iter.actual = iter.actual.prox
		iter.anterior = iter.anterior.prox
	}
}
func (lista *listaEnlazada[T]) Insertar(elemento T) {
	// Esta me tosqueó
}
func (lista *listaEnlazada[T]) Borrar() T {
	// Esta funcion medio que me dejo en conflicto, la voy a hacer poniendo panics si se hacen operaciones invalidas, te pido que le pegues una chequeada y hagas los cambios que consideres necesarios xxq dudo q esta bien de una, pero creo que es la base :$
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	//Si mi elemento es el primero
	elemento = iter.actual.dato
	if iter.actual == iter.anterior {
		// Hago que el que era el primero sea el segundo
		iter.actual = iter.actual.prox
	} else {
		iter.anterior.prox = iter.actual.prox
	}
	iter.actual = iter.actual.prox
	iter.anterior = iter.anterior.prox
	return elemento
}
