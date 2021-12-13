package main

import (
	"fmt"
	"os"
	"netstats"
)

func main(){
	if len(os.Args)!=2{
		fmt.Println("Cantidad de argumentos invalidos")
	}else{
		netstat := netstats.Crear()
		netstat.CargarArchivo(os.Args[1])
		netstat.EscucharComandos()
	}
}