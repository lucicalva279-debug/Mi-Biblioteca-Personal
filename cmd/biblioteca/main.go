// package main: Le dice a Go que este es el archivo principal,
// hace funcionar todo el programa.
package main

// Trae mis propias carpetas de código.
import "mi-biblioteca-personal/internal/menu"

func main() {
	menu.Iniciar()
}
