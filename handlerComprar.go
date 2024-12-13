
package main

import (
    "github.com/gin-gonic/gin"
)

func comprarHandler(c *gin.Context) {
    c.HTML(200, "comprar.html", nil)
}
