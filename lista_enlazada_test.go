package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	cantidad_volumen_medio  = 10
	cantidad_volumen_grande = 1000
)

func TestListaVacia(t *testing.T) {
	// Test operaciones basicas lista vacia
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 4, lista.VerUltimo())
}

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
	arreglo := make([]int, cantidad_volumen_medio)
	for i := 0; i < cantidad_volumen_medio; i++ {
		arreglo[i] = i
	}
	for _, elemento := range arreglo {
		lista.InsertarPrimero(elemento)
		require.EqualValues(t, elemento, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	for j := 0; j < cantidad_volumen_medio; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}
func TestVolumenGrande(t *testing.T) {
	// Test volumen grande
	lista := TDALista.CrearListaEnlazada[int]()
	arreglo := make([]int, cantidad_volumen_grande)
	for h := 0; h < cantidad_volumen_grande; h++ {
		arreglo[h] = h
	}
	for _, elemento := range arreglo {
		lista.InsertarPrimero(elemento)
		require.EqualValues(t, elemento, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	for j := 0; j < cantidad_volumen_grande; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}
