package contenidoWeb

import (
	"fmt"
)

type ContenidoWeb struct {
	Multimedia []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, m := range cw.Multimedia {
		m.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}


type Imagen struct {
	Titulo string
	Formato string
	Canales float64
}

type Audio struct {
	Titulo string
	Formato string
	Duracion float64
}

type Video struct {
	Titulo string
	Formato string
	Frames float64
}

func (i *Imagen) Mostrar() {
	fmt.Println("---")
	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canales: ", i.Canales)
	fmt.Println("----")
}

func (a *Audio) Mostrar() {
	fmt.Println("---")
	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duracion: ", a.Duracion)
	fmt.Println("----")
}

func (v *Video) Mostrar() {
	fmt.Println("---")
	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)
	fmt.Println("----")
}