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
func conectarDB() (*sql.DB, error) {

	// Variables de configuración
	var (
		dbUser     = "root"                // Usuario de la base de datos
		dbPassword = "root"                // Contraseña de la base de datos
		dbHost     = "localhost"           // Dirección del servidor MySQL
		dbPort     = "3306"                // Puerto del servidor MySQL
		dbName     = "peliculas_streaming" // Nombre de la base de datos
	)

	type Pelicula struct {
		ID            int     `json:"id"`
		Titulo        string  `json:"titulo"`
		Genero        string  `json:"genero"`
		Anio          int     `json:"anio"`
		Imagen        string  `json:"imagen"`
		Precio        float64 `json:"precio"`
		FechaCreacion string  `json:"fecha_creacion"`
	}
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
func apiPeliculasHandler(c *gin.Context) {
	type Pelicula struct {
		ID            int     `json:"id"`
		Titulo        string  `json:"titulo"`
		Genero        string  `json:"genero"`
		Anio          int     `json:"anio"`
		Imagen        string  `json:"imagen"`
		Precio        float64 `json:"precio"`
		FechaCreacion string  `json:"fecha_creacion"`
	}

	db, err := conectarDB()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	defer db.Close()

	// Consulta a la base de datos
	query := "SELECT id, titulo, genero, anio, imagen, precio, fecha_creacion FROM peliculas order by RAND()"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener las películas"})
		return
	}
	defer rows.Close()

	// Procesar los resultados
	var peliculas []Pelicula
	for rows.Next() {
		var p Pelicula
		if err := rows.Scan(&p.ID, &p.Titulo, &p.Genero, &p.Anio, &p.Imagen, &p.Precio, &p.FechaCreacion); err != nil {
			log.Println("Error al procesar los resultados:", err)
			continue
		}
		peliculas = append(peliculas, p)
	}

	// Respuesta en formato JSON
	c.JSON(http.StatusOK, peliculas)
}
