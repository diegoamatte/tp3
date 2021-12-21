package biblioteca

import (
	"github.com/algo2/tp3/grafo"
)

/*
* Funcion ejecutada sobre vertices adyacentes en el grafo
*/
type fnVertice func(v,w, extra interface{})

/*
 Implementa el recorrido BFS.
*/
func bfs(grafo *grafo.Grafo, origen interface{}, fn fnVertice, extra interface{}) (padres map[interface{}]interface{},orden map[interface{}]int) {
	cola := ColaCrear()
	visitados := make(map[interface{}]int)
	padres = make(map[interface{}]interface{})
	orden = make(map[interface{}]int)

	visitados[origen] = 1
	padres[origen] = nil
	orden[origen] = 0
	cola.Encolar(origen)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, ady := range (*grafo).ObtenerAdyacentes(v) {
			if fn != nil {
				fn(v,ady, extra)
			}
			if _, ok := visitados[ady]; !ok { //Si no fue visitado
				cola.Encolar(ady)
				orden[ady] = orden[v] + 1
				padres[ady] = v
				visitados[ady] = 1
			}
		}
	}
	return padres, orden
}
