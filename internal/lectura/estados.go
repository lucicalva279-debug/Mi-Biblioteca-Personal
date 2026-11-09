package lectura

import "fmt"

// Comprueba si el estado ingresado es válido.
func ValidarEstado(estado string) bool {
	// Retorna true solo si coincide con una de estas tres palabras.
	return estado == "Pendiente" ||
		estado == "Leyendo" ||
		estado == "Terminado"
}

// Cambia el estado de lectura de un libro.
func ActualizarEstado(estados []string, indice int, nuevoEstado string) error {
	// Evita buscar un cajón negativo o un número mayor a los libros que tenemos.
	if indice < 0 || indice >= len(estados) {
		return fmt.Errorf("el libro no existe")
	}
	// Si el estado es inventado, frena el código y devuelve error.
	if !ValidarEstado(nuevoEstado) {
		return fmt.Errorf("el estado ingresado no es válido")
	}
	//Como todo salió bien y sí se cambió, avisa que no hubo ningún error "nil".
	estados[indice] = nuevoEstado
	return nil
}

// Muestra los libros que tienen un estado determinado.
func ConsultarPorEstado(titulos, autores, estados []string, estado string) []string {
	var resultado []string
	// Recorre las listas posición por posición (0, 1, 2...).
	for i := 0; i < len(estados); i++ {
		// Si el estado del libro actual es igual al que estamos buscando...
		if estados[i] == estado {
			// Une el Título y el Autor con un guion y lo mete a la lista de resultados.
			resultado = append(resultado, titulos[i]+" - "+autores[i])
		}
	}
	return resultado // Devuelve la lista con todos los libros que coincidieron.
}

// Cuenta los libros que tienen un estado determinado.
func ContarPorEstado(estados []string, estado string) int {
	contador := 0
	// "for range" revisa cada elemento de la lista uno por uno de forma automática.
	// Ignoramos el índice con el guion bajo (_) porque nos importa el texto (estadoLibro).
	for _, estadoLibro := range estados {
		// Si el estado de este libro es el que buscamos, sumamos 1 al contador.
		if estadoLibro == estado {
			contador++
		}
	}
	return contador // Devuelve la cantidad de libros que coincidieron.
}
