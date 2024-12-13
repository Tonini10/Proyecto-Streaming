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
func conectarDBpel() (*sql.DB, error) {
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
func apiPeliculaHandler(c *gin.Context) {
	id := c.Param("id")
	type Pelicula struct {
		ID            int     `json:"id"`
		Titulo        string  `json:"titulo"`
		Genero        string  `json:"genero"`
		Anio          int     `json:"anio"`
		Imagen        string  `json:"imagen"`
		Precio        float64 `json:"precio"`
		FechaCreacion string  `json:"fecha_creacion"`
	}

	db, err := conectarDBpel()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	defer db.Close()

	// Consulta a la base de datos
	println(id)
	query := "SELECT id, titulo, genero, anio, imagen, precio, fecha_creacion FROM peliculas where id = " + id
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

// func estatico_apiPeliculasHandler(c *gin.Context) {
// 	peliculas := []Pelicula{
// 		{1, "El Padrino 2", "Drama", 1972, "/images/el_padrino.jpg"},
// 		{2, "Pulp Fiction", "Crimen", 1994, "/images/pulp_fiction.jpg"},
// 		{3, "El Caballero de la Noche", "Acción", 2008, "/images/el_caballero_de_la_noche.jpg"},
// 		{4, "Forrest Gump", "Drama", 1994, "/images/forrest_gump.jpg"},
// 		{5, "Inception", "Ciencia Ficción", 2010, "/images/inception.jpg"},
// 		{6, "Matrix", "Ciencia Ficción", 1999, "/images/matrix.jpg"},
// 		{7, "El Señor de los Anillos: La Comunidad del Anillo", "Fantasía", 2001, "/images/senor_de_los_anillos.jpg"},
// 		{8, "La Lista de Schindler", "Histórico", 1993, "/images/la_lista_de_schindler.jpg"},
// 		{9, "Gladiador", "Acción", 2000, "/images/gladiador.jpg"},
// 		{10, "Titanic", "Romance", 1997, "/images/titanic.jpg"},
// 		{11, "Jurassic Park", "Aventura", 1993, "/images/jurassic_park.jpg"},
// 		{12, "El Resplandor", "Terror", 1980, "/images/el_resplandor.jpg"},
// 		{13, "Los Infiltrados", "Crimen", 2006, "/images/los_infiltrados.jpg"},
// 		{14, "Memento", "Misterio", 2000, "/images/memento.jpg"},
// 		{15, "El Gran Lebowski", "Comedia", 1998, "/images/el_gran_lebowski.jpg"},
// 		{16, "Avatar", "Ciencia Ficción", 2009, "/images/avatar.jpg"},
// 		{17, "Amélie", "Romance", 2001, "/images/amelie.jpg"},
// 		{18, "Rocky", "Deporte", 1976, "/images/rocky.jpg"},
// 		{19, "Django Unchained", "Western", 2012, "/images/django_unchained.jpg"},
// 		{20, "La La Land", "Musical", 2016, "/images/la_la_land.jpg"},
// 		{21, "Harry Potter y la Piedra Filosofal", "Fantasía", 2001, "/images/harry_potter.jpg"},
// 		{22, "Casino Royale", "Acción", 2006, "/images/casino_royale.jpg"},
// 		{23, "El Origen", "Ciencia Ficción", 2010, "/images/el_origen.jpg"},
// 		{24, "Interestelar", "Ciencia Ficción", 2014, "/images/interestelar.jpg"},
// 		{25, "Toy Story", "Animación", 1995, "/images/toy_story.jpg"},
// 	}

// 	c.JSON(200, peliculas)
// }
