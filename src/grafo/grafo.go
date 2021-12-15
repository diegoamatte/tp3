package grafo

type Grafo interface {

	//Agrega un vertice al grafo. Si el vertice ya existia, no hace nada.
	//
	//Pre: el grafo debe estar creado.
	//
	//Post: se agrego un vertice sin aristas al grafo, si este no existia.
	AgregarVertice(v interface{})

	//Borra un vertice y todas sus aristas del grafo. Si el vertice no existe, no hace nada.
	//
	//Pre: el vertice debe existir en el grafo.
	SacarVertice(v interface{})

	//Agrega una arista entre los vertices y el peso, pasados como parametro.
	//
	//Pre: los vertices deben existir.
	AgregarArista(v interface{}, w interface{}, peso int) bool

	//Saca una arista del grafo
	SacarArista(v interface{}, w interface{})

	//Devuelve true si los vertices estan unidos por una arista, false en caso contrario.
	EstanUnidos(v interface{}, w interface{}) bool

	//Devuelve un slice conteniendo a todos los vertices del grafo.
	ObtenerVertices() []interface{}

	//Devuelve un slice con todos los vertices adyacentes al vertice pasado como parametro.
	ObtenerAdyacentes(v interface{}) []interface{}

	//Devuelve un vertice aleatorio del grafo.
	VerticeAleatorio() (interface{}, bool)

	//Devuelve el peso de una arista
	PesoArista(v1 interface{}, v2 interface{}) (int, bool)
}
