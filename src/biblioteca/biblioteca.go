package biblioteca

import "github.com/algo2/tp3/grafo"

type Biblioteca interface {
	//Devuelve un slice con el camino mas corto entre el origen y el destino.
	//Si no existe solucion devuelve false como segundo parametro
	CaminoMasCorto(grafo *grafo.Grafo, origen interface{}, destino interface{}) ([]interface{}, bool)

	PageRank(grafo *grafo.Grafo, n int)

	Diametro(grafo *grafo.Grafo)([]interface{}, int)
}
