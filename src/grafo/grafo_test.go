package grafo

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGrafoVacio(t *testing.T) {
	grafo := Crear(false)
	require.Empty(t, grafo.ObtenerVertices(), "El grafo deberia estar vacio")
	_, ok := grafo.VerticeAleatorio()
	require.False(t, ok, "Vertice Aleatorio deberia devolver false si no tiene vertices")
}

func TestGrafoVertices(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F", "G"}
	grafo := Crear(false)
	for _, v := range vertices {
		grafo.AgregarVertice(v)
	}
	require.NotEmpty(t, grafo.ObtenerVertices(), "Deberia agregarse el vertice")
	require.ElementsMatch(t, vertices, grafo.ObtenerVertices(), "Deberian coincidir los vertices con los agregados")
	grafo.SacarVertice("A")
	vertices = vertices[1:7]
	require.ElementsMatch(t, vertices, grafo.ObtenerVertices(), "Deberian coincidir los vertices con los agregados")
}

func TestGrafoAristas(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F", "G"}
	grafo := Crear(false)
	for _, v := range vertices {
		grafo.AgregarVertice(v)
	}
	aristas := []string{"B", "C", "D"}
	for _, a := range aristas {
		grafo.AgregarArista("A", a, 1)
	}
	require.ElementsMatch(t, grafo.ObtenerAdyacentes("A"), aristas, "Las aristas deberian coincidir")
	require.False(t, grafo.AgregarArista("A", "H", 1), "No se deberia poder agregar una arista con un vertice inexistente")
	require.True(t, grafo.EstanUnidos("A", "B"), "Los vertices A y B deberian estar unidos.")
	require.False(t, grafo.EstanUnidos("B", "C"), "Los vertices C y D no deberian estar unidos")
}

func TestGrafoDirigido(t *testing.T) {
	grafo := Crear(true)
	vertices := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, v := range vertices {
		grafo.AgregarVertice(v)
	}
	grafo.AgregarArista("A", "B", 2)

	require.False(t, grafo.EstanUnidos("B", "A"), "Los vertices no deberian estar unidos en un grafo dirigido")
	require.True(t, grafo.EstanUnidos("A", "B"), "Los vertices deberian estar unidos")
	peso, _ := grafo.PesoArista("A", "B")
	require.Equal(t, 2, peso, "El peso deberia coincidir con el agregado")

}

func TestVolumen(t *testing.T) {
	start := time.Now()
	largo := 10000
	cantAristas := 4
	grafo := Crear(false)
	vertices := make([]int, 0, largo)
	for i := 0; i < largo; i++ {
		grafo.AgregarVertice(i)
		vertices = append(vertices, i)
	}
	ok := true
	for i := 0; i < largo; i++ {
		for j := 0; j < cantAristas; j++ {
			peso := rand.Intn(10)
			ok = grafo.AgregarArista(i, rand.Intn(largo), peso)
			if !ok {
				break
			}
		}
		if !ok {
			break
		}
	}
	require.True(t, ok, "Se deberian poder agregar todas las aristas")
	require.ElementsMatch(t, vertices, grafo.ObtenerVertices(), "Los vertices deberian coincidir con los agregados")

	elapsed := time.Since(start)
	log.Printf("volumen tomo %s", elapsed)
}
