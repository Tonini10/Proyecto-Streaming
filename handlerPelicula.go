package main

import (
	"github.com/gin-gonic/gin"
)

func peliculaHandler(c *gin.Context) {
	id := c.Param("id")
	println("Que Pelicula?")
	println(id)

	// Usa "pelicula.html" como plantilla y pasa el ID como dato a la plantilla
	c.HTML(200, "pelicula.html", gin.H{
		"id": id, // Envías el ID a la plantilla
	})
}
