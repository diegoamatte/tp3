package grafo

import (
	"math/rand"
	"time"
)

type arista struct {
	peso    int
	destino interface{}
}

type grafoType struct {
	vertices map[interface{}][]interface{}
	dirigido bool
}

func Crear(dirigido bool) Grafo {
	grafo := grafoType{}
	grafo.vertices = make(map[interface{}][]interface{})
	grafo.dirigido = dirigido
	return &grafo
}

func (grafo *grafoType) AgregarVertice(v interface{}) {
	_, ok := grafo.vertices[v]
	if ok {
		return
	}
	grafo.vertices[v] = make([]interface{}, 0)
}

func (grafo *grafoType) SacarVertice(v interface{}) {
	delete(grafo.vertices, v)
}

func (grafo *grafoType) AgregarArista(v1 interface{}, v2 interface{}, peso int) bool {
	arista1, ok1 := grafo.vertices[v1]
	arista2, ok2 := grafo.vertices[v2]
	if !ok1 || !ok2 {
		return false
	}
	arista1 = append(arista1, arista{peso: peso, destino: v2})
	grafo.vertices[v1] = arista1
	if !grafo.dirigido {
		arista2 = append(arista2, arista{peso: peso, destino: v1})
		grafo.vertices[v2] = arista2
	}
	return true
}

func (grafo *grafoType) SacarArista(v1 interface{}, v2 interface{}) {
	if !grafo.dirigido {
		aristas2 := grafo.vertices[v2]
		grafo.vertices[v2] = borrar(aristas2, v1)
	}
	aristas1 := grafo.vertices[v1]
	grafo.vertices[v1] = borrar(aristas1, v2)
}

func (grafo *grafoType) EstanUnidos(v1 interface{}, v2 interface{}) bool {
	aristas1 := grafo.vertices[v1]
	ok1 := false
	ok2 := false
	if grafo.dirigido {
		for i := range aristas1 {
			if aristas1[i].(arista).destino == v2 {
				ok1 = true
				break
			}
		}
		return ok1
	}
	aristas2 := grafo.vertices[v2]
	for i := range aristas2 {
		if aristas2[i].(arista).destino == v1 {
			ok2 = true
			break
		}
	}
	return ok2
}

func (grafo *grafoType) ObtenerVertices() []interface{} {
	vertices := make([]interface{}, 0, len(grafo.vertices))
	for clave := range grafo.vertices {
		vertices = append(vertices, clave)
	}
	return vertices
}

func (grafo *grafoType) ObtenerAdyacentes(v interface{}) []interface{} {
	aristas := make([]interface{}, 0)
	for _, v := range grafo.vertices[v] {
		aristas = append(aristas, v.(arista).destino)
	}
	return aristas
}

func (grafo *grafoType) VerticeAleatorio() (interface{}, bool) {
	if len(grafo.vertices) == 0 {
		return nil, false
	}
	// Seed random
	rand.Seed(time.Now().UnixNano())
	indexes := rand.Perm(len(grafo.vertices))
	return grafo.ObtenerVertices()[indexes[0]], true
}

func (grafo *grafoType) PesoArista(v1 interface{}, v2 interface{}) (int, bool) {
	if !grafo.EstanUnidos(v1, v2) {
		return 0, false
	}
	peso := 0
	ok := false
	aristas := grafo.vertices[v1]
	for i := range aristas {
		if aristas[i].(arista).destino == v2 {
			peso = aristas[i].(arista).peso
			ok = true
		}
	}
	return peso, ok
}

//Funcion auxiliar de borrado sobre slices eficiente
func borrar(aristas []interface{}, destino interface{}) []interface{} {
	for i := range aristas {
		if aristas[i] == destino {
			aristas[i] = aristas[len(aristas)-1] // Se copia el ultimo elemento a la posicion a borrar
			aristas = aristas[:len(aristas)-1]
			break
		}

	}
	return aristas
}
