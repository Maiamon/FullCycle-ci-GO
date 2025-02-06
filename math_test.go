package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Somar(10, 5)
	if resultado != 15 {
		t.Errorf("Soma(10, 5) = %d; want 15", resultado)
	}
}
