package biblioteca

import (
	"github.com/algo2/tp3/grafo"
)


/*
 Implementa el recorrido BFS.
 Recibe por parametro una funcion que sera ejecutada sobre los diccionarios de padres y de orden y
 devolvera el resultado de aplicarles la funcion.
 Si recibe nil, devuelve dichos diccionarios
*/
func bfs(grafo *grafo.Grafo, origen interface{}) (map[interface{}]interface{}, map[interface{}]int) {
	cola := ColaCrear()
	visitados := make(map[interface{}]int)
	padres := make(map[interface{}]interface{})
	orden := make(map[interface{}]int)

	visitados[origen] = 1
	padres[origen] = nil
	orden[origen] = 0
	cola.Encolar(origen)

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, ady := range (*grafo).ObtenerAdyacentes(v) {
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
