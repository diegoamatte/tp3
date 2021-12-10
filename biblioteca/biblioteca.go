package biblioteca

type Biblioteca interface {
	//Devuelve un slice con el camino mas corto entre el origen y el destino.
	//Si no existe solucion devuelve false como segundo parametro
	CaminoMasCorto(origen interface{}, destino interface{}) ([]interface{}, bool)
}
