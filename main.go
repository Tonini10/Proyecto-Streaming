package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files (images)
	router.Static("/images", "./static/images")
	router.Static("/assets", "./static/assets/")
	router.Static("/css", "./templates/css")

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", mainPage)
	router.GET("/peliculas", peliculasHandler)
	router.GET("/pelicula/:id", peliculaHandler)
	router.GET("/crear-usuario", crearUsuarioHandler)
	router.GET("/comprar", comprarHandler)
	router.GET("/ver-compras", verComprasHandler)
	router.GET("/api/peliculas", apiPeliculasHandler) // API for movies
	router.GET("/api/pelicula/:id", apiPeliculaHandler)
	router.GET("/api/peliculayoutube/:id", apiPeliculaYotubeHandler)
	router.GET("/api/peliculasemail/:email", apiPeliculasEmailHandler) // API for movies
	router.GET("/api/peliculasemailtotal/:email", apiPeliculasEmailTotalHandler)
	router.GET("/api/usuarioemail/:email", apiUsuarioEmailHandler)
	router.POST("/api/compra", apiCompraHandler)
	router.GET("/micuenta/", miCuentaHandler)

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Cambia esto por el dominio permitido
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Run(":8080")
}
