package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	cantidadVolumenMedio  = 10
	cantidadVolumenGrande = 1000
)

func TestInsertaPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerPrimero())
	lista.InsertarPrimero(3)
	require.EqualValues(t, 3, lista.VerPrimero())

}

func TestInsertaUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	lista.InsertarUltimo(1.25)
	require.EqualValues(t, 1.25, lista.VerUltimo())
	lista.InsertarUltimo(2.96)
	require.EqualValues(t, 2.96, lista.VerUltimo())
	lista.InsertarUltimo(3.03)
	require.EqualValues(t, 3.03, lista.VerUltimo())

}

func TestBorroPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero("HOLA")
	lista.InsertarUltimo("COMO")
	lista.InsertarUltimo("ESTAS")
	require.EqualValues(t, "HOLA", lista.VerPrimero())
	require.EqualValues(t, "ESTAS", lista.VerUltimo())
	require.EqualValues(t, "HOLA", lista.BorrarPrimero())
	require.EqualValues(t, "COMO", lista.VerPrimero())
	require.EqualValues(t, "ESTAS", lista.VerUltimo())
	require.EqualValues(t, "COMO", lista.BorrarPrimero())
	require.EqualValues(t, "ESTAS", lista.VerPrimero())
	require.EqualValues(t, "ESTAS", lista.VerUltimo())
	require.EqualValues(t, "ESTAS", lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

// Test para comprobar funcion Largo
func TestLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.Largo())
	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero(3)
	require.EqualValues(t, 3, lista.Largo())
	lista.BorrarPrimero()
	require.EqualValues(t, 2, lista.Largo())
	lista.BorrarPrimero()
	require.EqualValues(t, 1, lista.Largo())
	lista.BorrarPrimero()
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

// Test operaciones basicas lista vacia
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 4, lista.VerUltimo())
}

// Test de volumen
func TestVolumenChico(t *testing.T) {
	//	Test volumen pequeño pero con datos de tipo string
	lista := TDALista.CrearListaEnlazada[string]()
	arreglo := []string{"tipo", "de", "dato", "abstracto"}
	for i := 0; i < 4; i++ {
		elemento := arreglo[i]
		lista.InsertarPrimero(elemento)
		require.EqualValues(t, elemento, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	for j := 0; j < 4; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}
func TestVolumenMedio(t *testing.T) {
	// Test volumen medio
	lista := TDALista.CrearListaEnlazada[int]()
	arreglo := make([]int, cantidadVolumenMedio)
	for i := 0; i < cantidadVolumenMedio; i++ {
		arreglo[i] = i
	}
	for _, elemento := range arreglo {
		lista.InsertarPrimero(elemento)
		require.EqualValues(t, elemento, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	for j := 0; j < cantidadVolumenMedio; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}
func TestVolumenGrande(t *testing.T) {
	// Test volumen grande
	lista := TDALista.CrearListaEnlazada[int]()
	arreglo := make([]int, cantidadVolumenGrande)
	for h := 0; h < cantidadVolumenGrande; h++ {
		arreglo[h] = h
	}
	for _, elemento := range arreglo {
		lista.InsertarPrimero(elemento)
		require.EqualValues(t, elemento, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	for j := 0; j < cantidadVolumenGrande; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}

// Casos del iterador externo:
/* func Test(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	/// Agregas elementos

	iter := lista.Iterador()
	// fot iter.Siguiente()...

} */

// Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
func TestIterInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(33)
	iter := lista.Iterador()
	elem := 77
	iter.Insertar(elem)
	require.EqualValues(t, elem, lista.VerPrimero())
}

// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
func TestIterInsetarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarUltimo(1.25)
	lista.InsertarUltimo(2.96)
	lista.InsertarUltimo(3.03)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	var dato float32 = 2.023
	iter.Insertar(dato)
	require.EqualValues(t, dato, lista.VerUltimo())
}

// Insertar un elemento en el medio se hace en la posición correcta.
func TestIterInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	mitad := lista.Largo() / 2
	elem := ("Barbara")
	arreglo_str := []string{"TDA", "Lista", "Enlazada", "Alan", "Grace"}
	// Agrego elementos a la lista
	for i := 0; i < len(arreglo_str); i++ {
		lista.InsertarPrimero(arreglo_str[i])
	}
	for i := 0; i < lista.Largo(); i++ {
		// Recorro la lista hasta largo/2 e inserto el elemento en la mitad de la misma
		if i == mitad {
			iter.Insertar(elem)
		}
	}
	for i := 0; i < mitad; i++ {
		iter.Siguiente()
	}
	require.EqualValues(t, elem, iter.VerActual())
}

// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func TestIterModificaPriemero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(33)
	lista.InsertarUltimo(77)
	// [33] -> [77]
	iter := lista.Iterador()
	iter.Borrar()
	// [77]
	require.EqualValues(t, 77, lista.VerPrimero())
}

// Remover el último elemento con el iterador cambia el último de la lista.
func TestIterRemoverFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarUltimo(1.25)
	lista.InsertarUltimo(2.96)
	lista.InsertarUltimo(3.03)
	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		contador++
		iter.Siguiente()
		if contador == 2 {
			iter.Borrar()
		}
	}
	require.EqualValues(t, 2.96, lista.VerUltimo())
}

// Borrar todos los elementos con el iterador
func TestIterBorrarTodosLosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	arreglo := make([]int, cantidadVolumenMedio)
	for i := 0; i < cantidadVolumenMedio; i++ {
		arreglo[i] = i
	}
	for iter.HaySiguiente() {
		iter.Borrar()
	}
	require.True(t, lista.EstaVacia())
}

// Verificar que al remover un elemento del medio, este no está.
func TestIterBorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	mitad := lista.Largo() / 2
	arregloInt := []int{1, 2, 3, 4, 5}
	// Agrego elementos a la lista
	for i := 0; i < len(arregloInt); i++ {
		lista.InsertarPrimero(arregloInt[i])
	}
	var elemento int
	contador := 0
	for iter.HaySiguiente() {
		// Recorro la lista hasta largo/2 e inserto el elemento en la mitad de la misma
		contador++
		if contador == mitad {
			elemento = iter.Borrar()
		}
	}
	for i := 0; i < mitad; i++ {
		iter.Siguiente()
		require.True(t, compararElementos(elemento, iter.VerActual()))
	}

}
func compararElementos(a, b int) bool {
	return a == b
}

//Casos del iterador interno

// Test suma elementos
func TestContarElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(8)

	contador := 0
	contadorPuentero := &contador
	lista.Iterar(func(v int) bool {
		*contadorPuentero += v
		return true
	})
	require.EqualValues(t, 14, contador)
}

func TestContarElementosCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(8)
	resultado := 0
	contador := 0
	lista.Iterar(func(v int) bool {
		contador++
		resultado += v
		return contador < 2
	})
	require.EqualValues(t, 5, resultado)
}

// // Test busqueda lineal
func TestIterBusquedaLineal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(8)
	elemBuscado := 8
	pos := 0
	encontrado := false
	encontradoPuntero := &encontrado
	lista.Iterar(func(n int) bool {
		if n == elemBuscado {
			*encontradoPuntero = true
			return false
		}
		pos++
		require.EqualValues(t, 4, pos)
		return true
	})
}
