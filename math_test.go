package main

func TestSoma(t *testing.T) {
	resultado := Soma(10, 5)
	if resultado != 15 {
		t.Errorf("Soma(10, 5) = %d; want 15", resultado)
	}
}