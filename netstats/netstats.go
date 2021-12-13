package netstats

type Netstats interface {
	// Carga un archivo en formato tsv y lo convierte en un grafo
	CargarArchivo(path string)
	//Pre: se debe haber cargado el archivo
	EscucharComandos()
}
