package biblioteca

import (
	"math/rand"
	"github.com/algo2/tp3/grafo"
)

//Funcion aplicada dentro del algoritmo de randomwalks
type rwAction func(g *grafo.Grafo,v interface{},extra interface{})
//Funcion a realizarse en el BFS
type bfsAccion func(g *grafo.Grafo,v interface{}, extra interface{})

/*
 Implementa el recorrido BFS.
 Recibe por parametro una funcion que sera ejecutada sobre los diccionarios de padres y de orden y
 devolvera el resultado de aplicarles la funcion.
 Si recibe nil, devuelve dichos diccionarios
*/
func bfs(grafo *grafo.Grafo, origen interface{}, fn bfsAccion,extra interface{}) (map[interface{}]interface{}, map[interface{}]int) {
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
				if fn != nil{
					fn(grafo,v,extra)
				}
			}
		}
	}
	return padres, orden
}

/*
* Algoritmo de Randomwalks
 */
func randomWalk(grafo *grafo.Grafo, largoRecorrido int, iteraciones int,fn rwAction, extra interface{}) {
	for i := 0; i < iteraciones; i++ {
		v, _ := (*grafo).VerticeAleatorio()
		for j := 0; j < largoRecorrido; j++ {
			vecino := vecinoAleatorio(grafo, v)
			fn(grafo,v,extra)
			v = vecino
		}
	}
}

/*

 */
func vecinoAleatorio(grafo *grafo.Grafo, vertice interface{}) interface{} {
	vecinos := (*grafo).ObtenerAdyacentes(vertice)
	cantVecinos := len(vecinos)
	var indice int
	if cantVecinos == 0 {
		return nil
	}
	indice = rand.Intn(cantVecinos)
	return vecinos[indice]
}
