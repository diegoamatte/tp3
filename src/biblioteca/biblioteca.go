package biblioteca

import "github.com/algo2/tp3/grafo"

type Biblioteca interface {
	//Devuelve un slice con el camino mas corto entre el origen y el destino.
	//Si no existe solucion devuelve false como segundo parametro
	CaminoMasCorto(grafo *grafo.Grafo, origen interface{}, destino interface{}) ([]interface{}, bool)

	PageRank(grafo *grafo.Grafo, n int)

	Diametro(grafo *grafo.Grafo) ([]interface{}, int)

	//Devuelve la cantidad de vertices que se encuentran a un rango n del origen
	Rango(grafo *grafo.Grafo, origen interface{}, n int) int

	Navegacion(grafo *grafo.Grafo, origen interface{}, n int)interface{}

	Conectividad(grafo *grafo.Grafo, pagina interface{})
}
