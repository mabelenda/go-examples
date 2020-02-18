package main

import "fmt"

type pasajero struct {
	tipoPasajero     string
	porcentajeAPagar int
}

//PORCENTAJEBASE PORCENTAJEMINUTO PORCENTAJEEMPLEADO son constantes que representan el porcentaje total a pagar
const (
	PORCENTAJEBASE     = 100
	PORCENTAJEMINUTO   = 50
	PORCENTAJEEMPLEADO = 0
)

func main() {

	pasajeros := []pasajero{
		pasajero{"base", PORCENTAJEBASE},
		pasajero{"ultimo Momento", PORCENTAJEMINUTO},
		pasajero{"Empleado", PORCENTAJEEMPLEADO},
		pasajero{"empleado de aerolinea de ultimo minuto", PORCENTAJEMINUTO + PORCENTAJEEMPLEADO},
	}

	var valorTicket int

	fmt.Print("Ingrese el precio base del Ticket para este vuelo:")
	fmt.Scanf("%v", &valorTicket)

	fmt.Println("LOS PRECIOS FINALES:")

	for _, v := range pasajeros {
		fmt.Println("Precio para el pasajero", v.tipoPasajero, "es", calcularPrecio(valorTicket, v.porcentajeAPagar))
	}
}

func calcularPrecio(precioTicket, porcentaje int) float64 {
	return float64((precioTicket * porcentaje) / 100)
}
