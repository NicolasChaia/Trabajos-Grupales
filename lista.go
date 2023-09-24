package lista
//Archivo con las primitivas
type Lista[T any] interface {
	// Corrobora si una lista esta vacia, devolviendo un dato del tipo booleano en tal caso.
	EstaVacia() bool

	// Permite insertar un elemento en la primera poscicion de una lista y en tiempo constante.
	InsertarPrimero(T)

	// Permite insertar un elemento en la ultima poscicio de la lista.
	InsertarUltimo(T)

	// Se encarga de borra el primer elemento de la lista enlazada y devolverlo.
	BorrarPrimero() T

	// Permite mostrar el dato del primer elemento de la lista.
	VerPrimero() T

	// Permite mostrar el dato del ultimo elemento de la lista.
	VerUltimo() T

	// Esta primitiva permite al usuario saber el largo de la lista.
	Largo() int

	// Permite aplicar la funcion a todos los elementos de la lista, teniendo dos condiciones de corte, que no haya mas elementos en la lista o que se devuelva falso.
	Iterar(visitar func(T) bool)

	// Devuelve un iterador propio de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// Permite ver el elemento actual que se esta iterando, si se quiere utilizar sobre un iterador que ya haya iterado todos los elementos, debe entrar en pánico con un mensaje El iterador termino de iterar.
	VerActual() T

	// Permite saber si hay un proximo elemento al actual.
	HaySiguiente() bool

	// Avanza al siguiente elemento de la lista. Si se quiere utilizar sobre un iterador que ya haya iterado todos los elementos, entra en pánico con un mensaje El iterador termino de iterar.
	Siguiente()

	// Inserta un elemento en la poscion anterior a la actual de iteracion, luego el iterador queda poscicionado sobre ese nuevo elemento insertado.
	Insertar(T)

	// Permite borrar, durante la iteracion, el elemento actual para luego devolverlo.
	Borrar() T
}
