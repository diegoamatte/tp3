package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Cantidad de argumentos invalidos")
	} else {
		netstat := Crear()
		netstat.CargarArchivo(os.Args[1])
		netstat.EscucharComandos()
	}
}
