package main

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"log"
	"net/http"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// conectarDB establece una conexión con la base de datos MySQL
func conectarDBUsuaEmail() (*sql.DB, error) {

	// Variables de configuración
	var (
		dbUser     = "root"                // Usuario de la base de datos
		dbPassword = "root"                // Contraseña de la base de datos
		dbHost     = "localhost"           // Dirección del servidor MySQL
		dbPort     = "3306"                // Puerto del servidor MySQL
		dbName     = "peliculas_streaming" // Nombre de la base de datos
	)

	// dsn := "usuario:contraseña@tcp(localhost:3306)/nombre_base_datos" // Cambia esto con tus credenciales
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

// apiPeliculasHandler maneja el endpoint para obtener películas
func apiUsuarioEmailHandler(c *gin.Context) {

	type Usuario struct {
		Nombre string `json:"nombre"`
	}

	db, err := conectarDBUsuaEmail()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	defer db.Close()

	email := c.Param("email")
	fmt.Println(email)

	// Consulta a la base de datos
	query := "SELECT CONCAT(nombre,' ',apellido) as nombre FROM `compras` where email = '" + email + "' limit 1 ;"
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener las películas" + email})
		return
	}
	defer rows.Close()

	// Procesar los resultados
	var usuario []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.Nombre); err != nil {
			log.Println("Error al procesar los resultados:", err)
			continue
		}
		usuario = append(usuario, u)
	}

	// Respuesta en formato JSON
	c.JSON(http.StatusOK, usuario)
}