package views

import (
	"estacionamiento/models"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	fondoEscena   *pixel.Sprite
	imagenFondo   pixel.Picture
	spriteVehiculo *pixel.Sprite
)

func cargarFondo() {
	file, err := os.Open("Assets/background.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	imagenFondo = pixel.PictureDataFromImage(img)
	fondoEscena = pixel.NewSprite(imagenFondo, imagenFondo.Bounds())
}

func cargarImagenVehiculo(ruta string) *pixel.Sprite {
	file, err := os.Open(ruta)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds())
}

func DibujarEscenaEstacionamiento(win *pixelgl.Window, vehiculos []models.Vehiculo) {
	if fondoEscena == nil {
		cargarFondo()
	}

	if spriteVehiculo == nil {
		spriteVehiculo = cargarImagenVehiculo("Assets/car.png")
	}

	fondoEscena.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	// Dibujar líneas de estacionamiento
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(2)
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.Line(2)
	imd.Push(pixel.V(700, 100), pixel.V(700, 500))
	imd.Line(2)

	anchoEstacionamiento := 600.0
	anchoCarril := anchoEstacionamiento / 10

	for i := 0.0; i < 10.0; i++ {
		offsetX := 100.0 + i*anchoCarril
		imd.Push(pixel.V(offsetX, 500), pixel.V(offsetX, 350))
		imd.Line(2)
		imd.Push(pixel.V(offsetX, 250), pixel.V(offsetX, 100))
		imd.Line(2)
	}

	// Dibujar vehículos
	for _, vehiculo := range vehiculos {
		spriteVehiculo.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.3).Moved(vehiculo.Posicion))
	}

	imd.Draw(win)
}