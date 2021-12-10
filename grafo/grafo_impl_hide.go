package grafo

import (
	"math/rand"
	"time"
)

type vertice struct {
	dato    interface{}
	aristas map[interface{}]int
}

type grafoType struct {
	vertices map[interface{}]vertice
	dirigido bool
}

func Crear(dirigido bool) Grafo {
	grafo := grafoType{}
	grafo.vertices = make(map[interface{}]vertice)
	grafo.dirigido = dirigido
	return &grafo
}

func (grafo *grafoType) AgregarVertice(v interface{}) {
	_, ok := grafo.vertices[v]
	if ok {
		return
	}
	nuevo_vertice := vertice{}
	nuevo_vertice.dato = v
	nuevo_vertice.aristas = make(map[interface{}]int)
	grafo.vertices[v] = nuevo_vertice
}

func (grafo *grafoType) SacarVertice(v interface{}) {
	delete(grafo.vertices, v)
}

func (grafo *grafoType) AgregarArista(v1 interface{}, v2 interface{}, peso int) bool {
	vertice1, ok1 := grafo.vertices[v1]
	vertice2, ok2 := grafo.vertices[v2]
	if !ok1 || !ok2 {
		return false
	}
	vertice1.aristas[v2] = peso
	if !grafo.dirigido {
		vertice2.aristas[v1] = peso
	}
	return true
}

func (grafo *grafoType) SacarArista(v1 interface{}, v2 interface{}) {
	if !grafo.dirigido {
		delete(grafo.vertices[v2].aristas, v1)
	}
	delete(grafo.vertices[v1].aristas, v2)
}

func (grafo *grafoType) EstanUnidos(v1 interface{}, v2 interface{}) bool {
	_, ok1 := grafo.vertices[v1].aristas[v2]
	if grafo.dirigido {
		return ok1
	}
	_, ok2 := grafo.vertices[v2].aristas[v1]
	return ok1 || ok2
}

func (grafo *grafoType) ObtenerVertices() []interface{} {
	vertices := make([]interface{}, 0, len(grafo.vertices))
	for _, v := range grafo.vertices {
		vertices = append(vertices, v.dato)
	}
	return vertices
}

func (grafo *grafoType) ObtenerAdyacentes(v interface{}) []interface{} {
	vertice := grafo.vertices[v]
	aristas := make([]interface{}, 0, len(vertice.aristas))
	for arista := range vertice.aristas {
		aristas = append(aristas, arista)
	}
	return aristas
}

func (grafo *grafoType) VerticeAleatorio() (interface{}, bool) {
	// Seed random
	if len(grafo.vertices) == 0 {
		return nil, false
	}
	rand.Seed(time.Now().UnixNano())
	indexes := rand.Perm(len(grafo.vertices))
	return grafo.ObtenerVertices()[indexes[0]], true
}

func (grafo *grafoType) PesoArista(v1 interface{}, v2 interface{}) (int, bool) {
	if !grafo.EstanUnidos(v1, v2) {
		return 0, false
	}
	peso, ok := grafo.vertices[v1].aristas[v2]
	return peso, ok
}
