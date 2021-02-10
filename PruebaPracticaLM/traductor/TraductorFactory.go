package traductor

import (
	"PruebaPracticaLM/traductor/implementacion"
	"fmt"
)

func GetTraductor(tipoTraductor string) (Traductor, error) {
	switch tipoTraductor {
	case "TEXTBINARY":
		return implementacion.NewTraductorTextoABinario(), nil
	default:
		return nil, fmt.Errorf("Tipo de traductor invalido")
	}
}
