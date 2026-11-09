# Mi Biblioteca Personal

Programa en Go, para consola, que sirve para llevar el registro de mis libros: guardar su título y autor y saber cuáles están pendientes, cuáles estoy leyendo y cuáles ya terminé.

Proyecto del curso Programación Orientada a Objetos (Go) – Aprendizaje Autónomo 1.

## ¿Qué hace?

- Registrar un libro (título y autor, empieza como "Pendiente")
- Mostrar todos los libros
- Buscar un libro por título
- Cambiar el estado de lectura (Pendiente, Leyendo, Terminado)
- Consultar libros por estado
- Contar libros por estado
- Eliminar un libro

Los datos se guardan solo en memoria: se pierden al cerrar el programa.

## Estructura

```
mi-biblioteca-personal/
├── go.mod
├── README.md
├── cmd/biblioteca/main.go     # package main: solo arranca el programa
└── internal/
    ├── catalogo/libros.go     # agregar, listar, buscar y eliminar libros
    ├── lectura/estados.go     # validar, cambiar, consultar y contar estados
    └── menu/menu.go           # menú interactivo en consola
```

No se usan paquetes de terceros, solo la biblioteca estándar de Go.

## Cómo ejecutarlo
Requiere Go instalado.

```
go build ./...
go run ./cmd/biblioteca
```

## Autora
Luciana Calva
Ingeniería en Ciberseguridad
