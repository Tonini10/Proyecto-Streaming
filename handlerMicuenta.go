package main

import (
	"github.com/gin-gonic/gin"
)

func miCuentaHandler(c *gin.Context) {
	c.HTML(200, "micuenta.html", nil)
}
