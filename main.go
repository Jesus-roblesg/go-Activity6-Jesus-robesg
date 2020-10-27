package main

import (
	"fmt"
	"./contenidoWeb"
)

func main()  {
	
	// Varias instancias 
	var op int
	var titulo string
	var formato string
	var data float64
	flag := 0
	cm := contenidoWeb.ContenidoWeb{Multimedia: []contenidoWeb.Multimedia{}}
	for flag != 1 {

		fmt.Println(" 1.- Agregar Imagen")
		fmt.Println(" 2.- Agregar Audio")
		fmt.Println(" 3.- Agregar Video")
		fmt.Println(" 4.- Mostrar Todo")
		fmt.Println(" 5.- Salir")

		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("Ingresa Titulo")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa Formato")
			fmt.Scan(&formato)
			fmt.Println("Ingresa no. canales")
			fmt.Scan(&data)

			imagen := contenidoWeb.Imagen{Titulo: titulo, Formato: formato,Canales: data}
			imagen.Mostrar()
			cm.Multimedia = append(cm.Multimedia, &imagen)

		case 2:
			fmt.Println("Ingresa Titulo")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa Formato")
			fmt.Scan(&formato)
			fmt.Println("Ingresa duracion")
			fmt.Scan(&data)

			audio := contenidoWeb.Audio{Titulo: titulo, Formato: formato, Duracion: data}
			audio.Mostrar()
			cm.Multimedia = append(cm.Multimedia, &audio)
		case 3:
			fmt.Println("Ingresa Titulo")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa Formato")
			fmt.Scan(&formato)
			fmt.Println("Ingresa no. frames")
			fmt.Scan(&data)

			video := contenidoWeb.Video{Titulo: titulo, Formato: formato,Frames: data}
			video.Mostrar()
			cm.Multimedia = append(cm.Multimedia, &video)
		case 4:
			cm.Mostrar()
			
		case 5:
			flag = 1
		default:
			fmt.Println("Error")
		}
	}
}