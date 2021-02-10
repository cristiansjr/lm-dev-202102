package implementacion

import (
	"reflect"
	"testing"
)

func TestNewTraductorBinarioATexto(t *testing.T) {
	tests := []struct {
		name string
		want *traductorTextoABinario
	}{
		{
			name: "Test NewTraductorTextoABinario",
			want: &traductorTextoABinario{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTraductorTextoABinario()
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTraductorTextoABinario() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraductorTextoABinario_Traducir(t *testing.T) {
	//Estructrura para el manejo de argumentos del test
	type args struct {
		valor string
	}
	tests := []struct{
		name string
		args args
		want string
	}{
		{
			name: "Test traductorTextoABinario-Traducir Exitoso",
			args: args{valor: "texto"},
			want: "11101001100101111100011101001101111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traductor := NewTraductorTextoABinario()
			got := traductor.Traducir(tt.args.valor)
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTraductorTextoABinario() = %v, want %v", got, tt.want)
			}
		})
	}
}