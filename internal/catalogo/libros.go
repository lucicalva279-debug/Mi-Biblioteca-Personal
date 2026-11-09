package catalogo

import "fmt"

// Agrega un libro al catálogo.
func AgregarLibro(titulos, autores, estados *[]string, titulo, autor string) error {
	// Validación: Evita que el usuario guarde textos vacíos "".
	if titulo == "" || autor == "" {
		return fmt.Errorf("el título y el autor son obligatorios")
	}
	// Punteros (*): Usamos el asterisco para modificar directamente las listas originales del main.
	*titulos = append(*titulos, titulo)
	*autores = append(*autores, autor)
	*estados = append(*estados, "Pendiente") // Todo libro nuevo empieza automáticamente como Pendiente.
	// Éxito: El libro se agregó correctamente.
	return nil
}

// Muestra todos los libros registrados.
func ListarLibros(titulos, autores, estados []string) []string {
	// Lista vacía para acumular los textos formateados.
	var resultado []string
	// Recorre todos los libros usando la longitud de la lista de títulos.
	for i := 0; i < len(titulos); i++ {
		// Sprintf: Une el Título, Autor y Estado en un solo texto bonito y ordenado.
		resultado = append(resultado,
			fmt.Sprintf("%s - %s - %s", titulos[i], autores[i], estados[i]))
	}
	return resultado // Devuelve la lista completa de libros formateados.
}

// Busca un libro por su título.
func BuscarLibro(titulos []string, titulo string) int {
	// Recorre la lista guardando la posición en "i" y el texto en "libro".
	for i, libro := range titulos {
		// Si encontramos el título exacto que busca el usuario...
		if libro == titulo {
			//Devuelve la posición del libro encontrado (índice).
			return i
		}
	}
	// Si termina el bucle y no lo encuentra, devuelve -1 (no existe).
	return -1
}

// Elimina un libro del catálogo.
func EliminarLibro(titulos, autores, estados *[]string, indice int) error {
	// Revisa que el cajón (índice) exista antes de borrar.
	if indice < 0 || indice >= len(*titulos) {
		return fmt.Errorf("el libro no existe")
	}
	// Truco de Go [:indice] y [indice+1:]: Corta la lista antes del libro y después del libro,
	// y los junta con los tres puntos (...), desapareciendo el elemento que estaba en medio.
	*titulos = append((*titulos)[:indice], (*titulos)[indice+1:]...)
	*autores = append((*autores)[:indice], (*autores)[indice+1:]...)
	*estados = append((*estados)[:indice], (*estados)[indice+1:]...)
	// Éxito: El libro se eliminó correctamente.
	return nil
}
