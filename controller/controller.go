package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}