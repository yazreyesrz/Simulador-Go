package models

import (
	"time"
)

func ManejarMovimientoVehiculos() {
	for i := len(Vehiculos) - 1; i >= 0; i-- {
		if Vehiculos[i].Posicion.X < 100 && Vehiculos[i].Carril == -1 && !Vehiculos[i].Entrando {
			Vehiculos[i].Posicion.X += 10
			if Vehiculos[i].Posicion.X > 100 {
				Vehiculos[i].Posicion.X = 100
			}
		} else if Vehiculos[i].Carril != -1 && !Vehiculos[i].Estacionado {
			var xObjetivo, yObjetivo float64
			anchoCarril := 600.0 / 10
			if Vehiculos[i].Carril < 10 {
				xObjetivo = 100.0 + float64(Vehiculos[i].Carril)*anchoCarril + anchoCarril/2
				yObjetivo = 400 + (500-350)/2
			} else {
				xObjetivo = 100.0 + float64(Vehiculos[i].Carril-10)*anchoCarril + anchoCarril/2
				yObjetivo = 100 + (250-100)/2
			}
			EstacionarVehiculo(&Vehiculos[i], xObjetivo, yObjetivo)
		}
	}
	ManejarSalidaVehiculos()
}

func VerificarTodosEstacionados() bool {
	todosEstacionados := true
	for _, vehiculo := range Vehiculos {
		if !vehiculo.Estacionado {
			todosEstacionados = false
			break
		}
	}
	return todosEstacionados
}

func ManejarSalidaVehiculos() {
	for i := len(Vehiculos) - 1; i >= 0; i-- {
		if Vehiculos[i].Estacionado && time.Now().After(Vehiculos[i].HoraSalida) && !Vehiculos[i].Entrando {
			if !Vehiculos[i].Teletransportando {
				Vehiculos[i].Teletransportando = true
				Vehiculos[i].TiempoInicioTeletransportacion = time.Now()
				Vehiculos[i].Posicion.X = 50
				Vehiculos[i].Posicion.Y = 400
			} else if time.Since(Vehiculos[i].TiempoInicioTeletransportacion) >= time.Millisecond*500 {
				ActualizarEstadoCarril(Vehiculos[i].Carril, false)
				RemoverVehiculo(i)
			}
		}
	}
}
