package traductor

type Traductor interface {
	Traducir(valor interface{}) interface{}
}