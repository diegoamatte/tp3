package netstats

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"grafo"
	"io"
	"os"
)

var comandos = []string{
	"listar_operaciones",
}

type netstatsType struct{
	grafo grafo.Grafo
}

func Crear() Netstats{
	return new(netstatsType)
}

func (netstats *netstatsType)CargarArchivo(path string){
	file, error := os.Open(path)
	defer file.Close()
	if error != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	data:= formatter('\t',file)
	netstats.crearGrafo(data)
	file.Close()
}

//Crea un grafo a partir de una matriz
func (netstats *netstatsType)crearGrafo(data [][]string){
	graph := grafo.Crear(true)
	for _, dato := range data {
		for i := range dato {
			graph.AgregarVertice(dato[i])
		}
		for j := 1; j < len(dato); j++ {
			graph.AgregarArista(dato[j-1], dato[j], 0)
		}
	}
	netstats.grafo = graph
}

// Funcion auxiliar para obtener una matriz con los datos a partir del archivo.
func formatter(delimiter rune,file *os.File)([][]string){
	reader := csv.NewReader(file)
	reader.Comma = delimiter//'\t'
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo.")
		return nil
	}
	return data
}

func (netstat *netstatsType)EscucharComandos(){
	for {
		reader := bufio.NewReader(os.Stdin)
		comando, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}
		ejecutar(comando)
	}
}

func ejecutar(comando string){
	switch comando {
	case "listar_operaciones\n":
		listarOperaciones()
	}
}


func listarOperaciones(){
	fmt.Println("Hola")
}