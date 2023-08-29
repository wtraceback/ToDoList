package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hello world",
        })
    })

    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/*.html")
    r.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "index",
            "test": "test",
        })
    })

    r.Run()
}