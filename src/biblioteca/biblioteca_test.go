package biblioteca

import (
	"testing"

	"github.com/algo2/tp3/grafo"
	//"github.com/stretchr/testify/require"
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
	//caminoCorrecto := [4]string{"A", "D", "E", "F"}
	//camino, ok := CaminoMasCorto(grafo, "A", "F")
	//require.True(t, ok, "El camino entre A y F deberia existir")
	//require.ElementsMatch(t, caminoCorrecto, camino)
	//_, ok = CaminoMasCorto(grafo, "D", "C")
	//require.False(t, ok, "El camino no deberia existir en un grafo dirigido")
}

func TestCaminosMinimosPalabras(t *testing.T) {
	grafo := grafo.Crear(true)
	vertices := [6]string{"America", "Europa", "Asia", "Oceania", "Africa"}
	for _, data := range vertices {
		grafo.AgregarVertice(data)
	}
	grafo.AgregarArista("America", "Europa", 0)
	grafo.AgregarArista("Europa", "Africa", 0)
	grafo.AgregarArista("Oceania", "America", 0)
	grafo.AgregarArista("America", "Oceania", 0)
	grafo.AgregarArista("Africa", "Asia", 0)
	grafo.AgregarArista("Asia", "Oceania", 0)

	//caminoCorrecto := [4]string{"Africa", "Asia", "Oceania", "America"}
	//camino, ok := CaminoMasCorto(grafo, "Africa", "America")
	//require.True(t, ok, "El camino entre Asia y America deberia existir")
	//require.ElementsMatch(t, caminoCorrecto, camino)
}

func TestPageRank(t *testing.T) {
	grafo := grafo.Crear(true)
	vertices := [6]string{"America", "Europa", "Asia", "Oceania", "Africa"}
	for _, data := range vertices {
		grafo.AgregarVertice(data)
	}
	grafo.AgregarArista("America", "Oceania", 0)
	grafo.AgregarArista("America", "Europa", 0)
	grafo.AgregarArista("Asia", "Oceania", 0)
	grafo.AgregarArista("Oceania", "America", 0)
	grafo.AgregarArista("Europa", "Africa", 0)
	grafo.AgregarArista("Africa", "Asia", 0)
	grafo.AgregarArista("Africa", "Oceania", 0)
	grafo.AgregarArista("Asia", "America", 0)

	PageRank(&grafo, 2)
}

