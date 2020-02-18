package main

import "testing"

func TestCalcularPrecio(t *testing.T) {
	type prueba struct {
		pasajero    pasajero
		valorTicket int
		respuesta   float64
	}

	pruebas := []prueba{
		prueba{pasajero{"base", 100}, 100, 100},
		prueba{pasajero{"ultimo Momento", 50}, 100, 50},
		prueba{pasajero{"Empleado", 0}, 100, 0},
	}

	for _, v := range pruebas {

		result := calcularPrecio(v.valorTicket, v.pasajero.porcentajeAPagar)

		if result != v.respuesta {
			t.Error("Expected", v.respuesta, "Got", result)
		}
	}

}
