package models

import (
	"math/rand"
	"sync"
	"time"

	"github.com/faiface/pixel"
)

type Vehiculo struct {
	ID               int
	Posicion         pixel.Vec
	PosicionAnterior pixel.Vec
	Carril           int
	Estacionado      bool
	HoraSalida       time.Time
	Entrando         bool
	Teletransportando bool
	TiempoInicioTeletransportacion time.Time
}


//sujeto
var (
	CanalVehiculos  chan Vehiculo
	Vehiculos       []Vehiculo
	CandadoVehiculos sync.Mutex
)

func InicializarVehiculos() {
	CanalVehiculos = make(chan Vehiculo)
	go GenerarVehiculos()
}

func CrearVehiculo(id int) Vehiculo {
	CandadoVehiculos.Lock()
	defer CandadoVehiculos.Unlock()
	vehiculo := Vehiculo{
		ID:           id,
		Posicion:     pixel.V(0, 300),
		Carril:       -1,
		Estacionado:  false,
	}
	Vehiculos = append(Vehiculos, vehiculo)
	return vehiculo
}

func AsignarHoraSalida(vehiculo *Vehiculo) {
	rand.Seed(time.Now().UnixNano())
	tiempoSalida := time.Duration(rand.Intn(10)+9) * time.Second
	vehiculo.HoraSalida = time.Now().Add(tiempoSalida)
}

func ObtenerVehiculos() []Vehiculo {
	return Vehiculos
}

func AsignarCarrilAVehiculo(id int, carril int) {
	CandadoVehiculos.Lock()
	defer CandadoVehiculos.Unlock()
	for i := range Vehiculos {
		if Vehiculos[i].ID == id {
			Vehiculos[i].Carril = carril
		}
	}
}

func ReiniciarPosicionVehiculo(id int) {
	CandadoVehiculos.Lock()
	defer CandadoVehiculos.Unlock()
	for i := range Vehiculos {
		if Vehiculos[i].ID == id {
			Vehiculos[i].Posicion = pixel.V(0, 300)
		}
	}
}

func EncontrarPosicionVehiculo(id int) pixel.Vec {
	CandadoVehiculos.Lock()
	defer CandadoVehiculos.Unlock()
	for _, vehiculo := range Vehiculos {
		if vehiculo.ID == id {
			return vehiculo.Posicion
		}
	}
	return pixel.Vec{}
}

func EstacionarVehiculo(vehiculo *Vehiculo, xObjetivo, yObjetivo float64) {
	vehiculo.Posicion.X = xObjetivo
	vehiculo.Posicion.Y = yObjetivo
	vehiculo.Estacionado = true
	AsignarHoraSalida(vehiculo)
}

func RemoverVehiculo(indice int) {
	Vehiculos = append(Vehiculos[:indice], Vehiculos[indice+1:]...)
}

func GenerarVehiculos() {
	id := 0
	for {
		id++
		vehiculo := CrearVehiculo(id)
		CanalVehiculos <- vehiculo
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	}
}
