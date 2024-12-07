package scenes

import (
	"estacionamiento/models"
	"estacionamiento/views"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func EjecutarSimulacion() {
	models.InicializarVehiculos()

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Simulador de Estacionamiento",
		Bounds: pixel.R(0, 0, 800, 600),
	})
	if err != nil {
		panic(err)
	}

	// Lanzar la goroutine que procesa los vehículos del canal

	//observador
	go ProcesarCanalVehiculos()

	for !win.Closed() {
		win.Clear(colornames.White)
		views.DibujarEscenaEstacionamiento(win, models.ObtenerVehiculos())
		win.Update()
		models.CandadoVehiculos.Lock()
		models.ManejarMovimientoVehiculos()
		models.CandadoVehiculos.Unlock()

		time.Sleep(16 * time.Millisecond)
	}
}  
  
      
func ProcesarCanalVehiculos() {
	for vehiculo := range models.CanalVehiculos {
		models.CandadoCarriles.Lock()
		for _, ocupado := range models.EstadoCarriles {
			if !ocupado {
				break
			}
		}
		models.CandadoCarriles.Unlock()
		models.GestionarCarrilVehiculo(vehiculo.ID)
	}
}