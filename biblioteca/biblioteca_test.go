package biblioteca

import (
	"github.com/stretchr/testify/require"
	"grafo"
	"testing"
)

func TestCaminosMinimos(t *testing.T) {
	grafo := grafo.Crear(true)
	vertices := [6]string{"A", "B", "C", "D", "E", "F"}
	for _, data := range vertices {
		grafo.AgregarVertice(data)
	}
	grafo.AgregarArista("A", "B", 0)
	grafo.AgregarArista("B", "C", 0)
	grafo.AgregarArista("C", "E", 0)
	grafo.AgregarArista("A", "D", 0)
	grafo.AgregarArista("D", "E", 0)
	grafo.AgregarArista("E", "F", 0)
	biblio := Crear(grafo)
	caminoCorrecto := [4]string{"A", "D", "E", "F"}
	camino, ok := biblio.CaminoMasCorto("A", "F")
	require.True(t, ok, "El camino entre A y F deberia existir")
	require.ElementsMatch(t, caminoCorrecto, camino)
	_, ok = biblio.CaminoMasCorto("D", "C")
	require.False(t, ok, "El camino no deberia existir en un grafo dirigido")
}
