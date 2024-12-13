
package main

import (
    "github.com/gin-gonic/gin"
)

func verComprasHandler(c *gin.Context) {
    c.HTML(200, "ver_compras.html", nil)
}
