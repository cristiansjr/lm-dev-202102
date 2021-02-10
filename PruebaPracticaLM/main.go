package main

import (
	"PruebaPracticaLM/traductor"
	"fmt"
)

func main() {
	textoATraducir := "texto"
	formatoOrigen := "TEXT"
	formatoDestino := "BINARY"
	textoTraducido := traducirTexto(textoATraducir,formatoOrigen, formatoDestino)
	valTexto := textoTraducido.(string)
	fmt.Println(valTexto)
}

func traducirTexto(textoATraducir string, formatoOrigen string, formatoDestino string) interface{} {
	tipoTraductor := formatoOrigen+formatoDestino
	traductorImp, err := traductor.GetTraductor(tipoTraductor)
	if err != nil {
		return ""
	}
	textoTraducido := traductorImp.Traducir(textoATraducir)
	return textoTraducido
}
