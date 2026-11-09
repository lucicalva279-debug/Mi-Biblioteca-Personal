// package menu: Contiene la lógica del menú interactivo de la biblioteca.
package menu
// import (): Es la lista de "herramientas" que vamos a pedir prestadas.
// Trae cosas de Go (como escribir en pantalla) y tus propias carpetas de código.
import (
    "bufio"   // Sirve para leer lo que el usuario escribe en el teclado.
    "fmt"     // Sirve para mostrar textos y mensajes en la pantalla.
    "os"      // Sirve para interactuar con el sistema operativo, como leer la entrada estándar.
    "strconv" // Sirve para transformar texto en números enteros.
    "strings" // Sirve para manipular cadenas de texto.
    // Trae las funciones de tu paquete "catalogo" (agregar, buscar, borrar libros).
    "mi-biblioteca-personal/internal/catalogo"
    // Trae las funciones de tu paquete "lectura" (cambiar estado, contar, validar estados).
    "mi-biblioteca-personal/internal/lectura"
)
func Iniciar() {
	// Listas principales en memoria que guardarán toda la base de datos de libros.
    var titulos []string
    var autores []string
    var estados []string
	// bufio: Herramienta estándar de Go que lee las entradas por teclado del usuario (os.Stdin).
    scanner := bufio.NewScanner(os.Stdin)
	// for sin condiciones crea un bucle infinito (el menú se repetirá hasta que elijas la opción 8).
    for {
        fmt.Println("\n--- MI BIBLIOTECA PERSONAL ---")
        fmt.Println("1. Registrar libro")
        fmt.Println("2. Mostrar libros")
        fmt.Println("3. Buscar libro")
        fmt.Println("4. Cambiar estado de lectura")
        fmt.Println("5. Consultar libros por estado")
        fmt.Println("6. Contar libros por estado")
        fmt.Println("7. Eliminar libro")
        fmt.Println("8. Salir")
        fmt.Print("Seleccione una opción: ")
        // leerTexto: Función auxiliar que lee una línea de texto del usuario.
        opcionTexto := leerTexto(scanner)
        // Convierte la opción ingresada por el usuario a un número entero.
        opcion, err := strconv.Atoi(opcionTexto)
         // Si el usuario escribe letras en lugar de números, se activa esta alerta.
        if err != nil {
            fmt.Println("Debe ingresar un número.")
			// Saltarse el resto de este turno del bucle y volver a mostrar el menú.
            continue 
        }
        // switch: Elige el camino correcto según el número que marcó el usuario.
        switch opcion {
        case 1:
            fmt.Print("Título: ")
            titulo := leerTexto(scanner)

            fmt.Print("Autor: ")
            autor := leerTexto(scanner)
            // El signo ampersand (&): Pasa las direcciones de memoria de las listas
            // para que AgregarLibro pueda inyectarles los nuevos datos definitivamente.
            err := catalogo.AgregarLibro(
                &titulos, &autores, &estados, titulo, autor,
            )

            if err != nil {
                fmt.Println(err) // Muestra el error si falta el título o autor.
            } else {
                fmt.Println("Libro registrado correctamente.")
            }

        case 2:
			// Obtiene la bolsa de textos formateados y la pasa a la función de imprimir.
            libros := catalogo.ListarLibros(titulos, autores, estados)
            mostrarResultados(libros)

        case 3:
            fmt.Print("Título que desea buscar: ")
            titulo := leerTexto(scanner)
            // Devuelve el número de cajón o índice donde se encuentra.
            indice := catalogo.BuscarLibro(titulos, titulo)

            if indice == -1 {
                fmt.Println("No se encontró el libro.")
            } else {
				// Si existe, usa ese índice para extraer la información exacta de las 3 listas.
                fmt.Println("Libro encontrado:", titulos[indice])
                fmt.Println("Autor:", autores[indice])
                fmt.Println("Estado:", estados[indice])
            }

        case 4:
            fmt.Print("Título del libro: ")
            titulo := leerTexto(scanner)
             // Paso 1: Busca en qué cajón está el libro.
            indice := catalogo.BuscarLibro(titulos, titulo)

            if indice == -1 {
                fmt.Println("No se encontró el libro.")
                continue
            }

            fmt.Print("Nuevo estado (Pendiente, Leyendo, Terminado): ")
            nuevoEstado := leerTexto(scanner)
            // Paso 2: Envía el índice encontrado a lectura para procesar el cambio.
            err := lectura.ActualizarEstado(estados, indice, nuevoEstado)

            if err != nil {
                fmt.Println(err) // Muestra error si la palabra escrita no es válida.
            } else {
                fmt.Println("Estado actualizado correctamente.")
            }

        case 5:
            fmt.Print("Estado que desea consultar: ")
			// El usuario escribe el estado que quiere consultar.
            estado := leerTexto(scanner)
            // Valida que el estado ingresado sea uno de los permitidos.
            if !lectura.ValidarEstado(estado) {
                fmt.Println("Estado no válido.")
                continue
            }
            // Llama a tu paquete 'lectura' para buscar y juntar en una lista 
            // todos los títulos y autores que tengan ese estado exacto.
            libros := lectura.ConsultarPorEstado(
                titulos, autores, estados, estado,
            )

            mostrarResultados(libros)

        case 6:
            fmt.Print("Estado que desea contar: ")
            estado := leerTexto(scanner)
            // Revisa que la palabra sea "Pendiente", "Leyendo" o "Terminado".
            if !lectura.ValidarEstado(estado) {
                fmt.Println("Estado no válido.")
                continue
            }
            // Llama a tu función para que recorra la lista y cuente (1, 2, 3...) cuántos coinciden
            cantidad := lectura.ContarPorEstado(estados, estado)
			// Muestra en la pantalla el número final (por ejemplo: "Cantidad de libros: 4").
            fmt.Println("Cantidad de libros:", cantidad)

        case 7:
            fmt.Print("Título del libro que desea eliminar: ")
            titulo := leerTexto(scanner)
            // Primero buscamos el índice (el cajón) usando el título proporcionado.
            indice := catalogo.BuscarLibro(titulos, titulo)

            // Si el índice es válido, procedemos a eliminar el libro.
			// Le mandamos las listas con el signo "&" para recortarlas en memoria permanentemente.
            err := catalogo.EliminarLibro(
                &titulos, &autores, &estados, indice,
            )

            if err != nil {
                fmt.Println(err) // Saltará error si "BuscarLibro" devolvió un -1.
            } else {
                fmt.Println("Libro eliminado correctamente.")
            }

        case 8:
			// Finaliza por completo la función main y apaga el programa de forma limpia.
            fmt.Println("Gracias por utilizar Mi Biblioteca Personal.")
            return

        default:
            fmt.Println("Opción no válida. Intente nuevamente.")
        }
    }
}
func leerTexto(scanner *bufio.Scanner) string {
    scanner.Scan() // Espera que el usuario escriba algo y presione Enter.
    return strings.TrimSpace(scanner.Text()) // Recorta espacios en blanco innecesarios.
}
// Imprime los elementos de una lista, o un mensaje si está vacía.
func mostrarResultados(resultados []string) {
	// Verifica si la lista de resultados está vacía.
    if len(resultados) == 0 { 
        fmt.Println("No hay libros para mostrar.")
        return // Corta la ejecución para que no intente recorrer nada abajo.
    }
    // Recorre e imprime cada línea formateada una por una.
    for _, resultado := range resultados {
        fmt.Println(resultado)
    }
}