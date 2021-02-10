package implementacion

type traductorBinarioATexto struct {
}

func NewTraductorBinarioATexto() *traductorBinarioATexto {
	return &traductorBinarioATexto{}
}

func (t *traductorBinarioATexto) Traducir(valor interface{}) interface{} {
	req := valor.(string)
	var normalString string

	normalString = req
	return normalString
}
