package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura que representa una compra
type Compra struct {
	IDCompra               int     `json:"id_compra"`
	Nombre                 string  `json:"nombre"`
	Apellido               string  `json:"apellido"`
	Email                  string  `json:"email"`
	NombreTarjeta          string  `json:"nombre_tarjeta"`
	NumeroTarjeta          string  `json:"numero_tarjeta"`
	FechaExpiracionTarjeta string  `json:"fecha_expiracion_tarjeta"`
	CVVTarjeta             string  `json:"cvv_tarjeta"`
	FechaCreacion          string  `json:"fecha_creacion"`
	IDPelicula             int     `json:"id_pelicula"`
	Costo                  float64 `json:"costo"`
}

// conectarDBpel establece una conexión con la base de datos MySQL
func conectarDBco() (*sql.DB, error) {
	// Variables de configuración
	var (
		dbUser     = "root"                // Usuario de la base de datos
		dbPassword = "root"                // Contraseña de la base de datos
		dbHost     = "localhost"           // Dirección del servidor MySQL
		dbPort     = "3306"                // Puerto del servidor MySQL
		dbName     = "peliculas_streaming" // Nombre de la base de datos
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// apiCompraHandler maneja el endpoint para registrar una compra
func apiCompraHandler(c *gin.Context) {
	// Crear una nueva instancia de la estructura Compra
	var compra Compra

	nombre := c.DefaultPostForm("nombre", "")
	apellido := c.DefaultPostForm("apellido", "")
	email := c.DefaultPostForm("email", "")
	nombreTarjeta := c.DefaultPostForm("nombre_tarjeta", "")
	numeroTarjeta := c.DefaultPostForm("numero_tarjeta", "")
	fechaExpiracionTarjeta := c.DefaultPostForm("fecha_expiracion_tarjeta", "")
	cvvTarjeta := c.DefaultPostForm("cvv_tarjeta", "")
	// Parse additional parameters: id_pelicula and costo
	idPelicula, err := strconv.Atoi(c.DefaultPostForm("id_pelicula", "0")) // Convert string to int, default to 0 if not provided
	if err != nil {
		// Handle the error (optional)
		log.Printf("Error converting id_pelicula: %v", err)
	}

	costo, err := strconv.ParseFloat(c.DefaultPostForm("costo", "0.0"), 64) // Convert string to float64, default to 0.0 if not provided
	if err != nil {
		// Handle the error (optional)
		log.Printf("Error converting costo: %v", err)
	}

	fmt.Println(nombre, apellido, email, nombreTarjeta, numeroTarjeta, fechaExpiracionTarjeta, cvvTarjeta, idPelicula, costo)

	// Validar que los campos necesarios estén presentes
	if nombre == "" || apellido == "" || email == "" || numeroTarjeta == "" || fechaExpiracionTarjeta == "" || cvvTarjeta == "" || idPelicula == 0 || costo == 0.0 {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		c.HTML(200, "comprar_error.html", nil)
		return
	}

	// Conectar a la base de datos
	db, err := conectarDBco()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		c.HTML(200, "comprar_error.html", nil)
		return
	}
	defer db.Close()

	// Insertar la compra en la base de datos
	query := `INSERT INTO compras (nombre, apellido, email, nombre_tarjeta, numero_tarjeta, fecha_expiracion_tarjeta, cvv_tarjeta, id_pelicula, costo) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, nombre, apellido, email, nombreTarjeta, numeroTarjeta, fechaExpiracionTarjeta, cvvTarjeta, idPelicula, costo)
	if err != nil {
		log.Println("Error al registrar la compra:", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar la compra"})
		c.HTML(200, "comprar_error.html", nil)
		return
	}

	// Obtener el ID de la nueva compra insertada
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID de la compra:", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el ID de la compra"})
		c.HTML(200, "comprar_error.html", nil)
		return
	}

	// Asignar el ID y fecha de creación a la compra
	compra.IDCompra = int(id)
	compra.Nombre = string(nombre)
	compra.Apellido = string(apellido)
	compra.Email = string(email)
	compra.NombreTarjeta = string(nombreTarjeta)
	compra.NumeroTarjeta = string(numeroTarjeta)
	compra.FechaExpiracionTarjeta = string(fechaExpiracionTarjeta)
	compra.CVVTarjeta = string(cvvTarjeta)
	// compra.FechaCreacion = string(f)
	compra.IDPelicula = int(idPelicula)
	compra.Costo = float64(costo)
	fmt.Println(compra)

	// Responder con los datos de la compra registrada
	// c.JSON(http.StatusOK, compra)
	c.HTML(200, "comprar.html", nil)
}
