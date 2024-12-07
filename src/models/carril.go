package models

import (
	"math/rand"
	"sync"
	"time"
)

const (
	numCarriles = 20
)

var (
	EstadoCarriles [numCarriles]bool
	CandadoCarriles sync.Mutex
)

func EsperarPosicionVehiculo(id int, xObjetivo float64) {
	for {
		posicion := EncontrarPosicionVehiculo(id)
		if posicion.X >= xObjetivo {
			break
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func BuscarCarrilDisponible() (int, bool) {
	CandadoCarriles.Lock()
	defer CandadoCarriles.Unlock()
	rand.Seed(time.Now().UnixNano())
	carriles := rand.Perm(numCarriles)
	for _, carril := range carriles {
		if !EstadoCarriles[carril] {
			EstadoCarriles[carril] = true
			return carril, true
		}
	}
	return -1, false
}

func GestionarCarrilVehiculo(id int) {
	CrearVehiculo(id)
	EsperarPosicionVehiculo(id, 100)
	carril, encontrado := BuscarCarrilDisponible()
	if !encontrado {
		ReiniciarPosicionVehiculo(id)
		return
	}
	AsignarCarrilAVehiculo(id, carril)
}

func ActualizarEstadoCarril(carril int, estado bool) {
	CandadoCarriles.Lock()
	defer CandadoCarriles.Unlock()
	EstadoCarriles[carril] = estado
}
