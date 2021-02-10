package implementacion

import "fmt"

type traductorTextoABinario struct {
}

func NewTraductorTextoABinario() *traductorTextoABinario {
	return &traductorTextoABinario{}
}

func (t *traductorTextoABinario) Traducir(valor interface{}) interface{} {
	req := valor.(string)
	var binString string
	for _, normalCharacter := range req {
		binString = fmt.Sprintf("%s%b", binString, normalCharacter)
	}
	return binString
}
