
package main

import (
    "github.com/gin-gonic/gin"
)

func crearUsuarioHandler(c *gin.Context) {
    c.HTML(200, "crear_usuario.html", nil)
}
