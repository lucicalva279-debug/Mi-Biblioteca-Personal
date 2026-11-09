// package catalogo: Contiene la lógica para gestionar el catálogo de libros.
package catalogo

// import: Trae la herramienta "fmt" de Go.
// Aquí la usamos para crear mensajes de error y textos con formato.
import "fmt"

// Agrega un libro al catálogo.
// Recibe punteros (*) a las tres listas para poder modificarlas de verdad,
// y recibe el título y el autor del libro nuevo. Devuelve un error si algo falla.
func AgregarLibro(titulos, autores, estados *[]string, titulo, autor string) error {
	// Validación: si el título o el autor están vacíos ("")...
	if titulo == "" || autor == "" {
		// ...frena la función y devuelve un error explicando el problema.
		return fmt.Errorf("el título y el autor son obligatorios")
	}
	// El asterisco (*) entra a la lista original. append agrega el título al final.
	*titulos = append(*titulos, titulo)
	// Agrega el autor en la misma posición que el título.
	*autores = append(*autores, autor)
	// Todo libro nuevo empieza con el estado "Pendiente".
	*estados = append(*estados, "Pendiente")
	// nil significa "no hubo ningún error": el libro se agregó correctamente.
	return nil
}

// Muestra todos los libros registrados.
// Recibe las tres listas y devuelve una lista de textos ya formateados.
func ListarLibros(titulos, autores, estados []string) []string {
	// Lista vacía donde iremos guardando el texto de cada libro.
	var resultado []string
	// Recorre las posiciones 0, 1, 2... hasta el total de títulos.
	for i := 0; i < len(titulos); i++ {
		// Agrega un texto nuevo a la lista de resultados.
		resultado = append(resultado,
			// Sprintf une título, autor y estado de la posición i en un solo texto.
			fmt.Sprintf("%s - %s - %s", titulos[i], autores[i], estados[i]))
	}
	// Devuelve la lista completa de libros formateados.
	return resultado
}

// Busca un libro por su título.
// Devuelve la posición (índice) donde está, o -1 si no lo encuentra.
func BuscarLibro(titulos []string, titulo string) int {
	// Recorre la lista: "i" es la posición y "libro" es el título de esa posición.
	for i, libro := range titulos {
		// Si el título de esta posición es exactamente el que busca el usuario...
		if libro == titulo {
			// ...devuelve la posición donde lo encontró.
			return i
		}
	}
	// Si terminó de recorrer toda la lista y no lo encontró, devuelve -1 (no existe).
	return -1
}

// Elimina un libro del catálogo.
// Recibe punteros a las tres listas y la posición (índice) del libro a borrar.
func EliminarLibro(titulos, autores, estados *[]string, indice int) error {
	// Revisa que la posición exista: no puede ser negativa ni mayor o igual a la cantidad de libros.
	if indice < 0 || indice >= len(*titulos) {
		// Si no existe, devuelve un error y no borra nada.
		return fmt.Errorf("el libro no existe")
	}
	// Truco de Go: [:indice] toma lo que está antes del libro y [indice+1:] lo que está después.
	// Con los tres puntos (...) se juntan las dos partes y el libro del medio desaparece.
	*titulos = append((*titulos)[:indice], (*titulos)[indice+1:]...)
	// Se borra la misma posición en la lista de autores para que no se mezclen los datos.
	*autores = append((*autores)[:indice], (*autores)[indice+1:]...)
	// Y también en la lista de estados.
	*estados = append((*estados)[:indice], (*estados)[indice+1:]...)

	// nil significa "no hubo ningún error": el libro se eliminó correctamente.
	return nil
}
