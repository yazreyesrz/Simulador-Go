package main

import (
	"estacionamiento/scenes"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(scenes.EjecutarSimulacion)
}
