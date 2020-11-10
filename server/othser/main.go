package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"openvpn-system/router"
)

func main() {
	fmt.Println("Blingabc BI_Platef starting ...")
	Rcontext := gin.Default()
	Rcontext.LoadHTMLGlob("templates/*")
	Rcontext.Static("/assets", "./assets")
	router.Distribute(Rcontext) //事项分支...

	// Rcontext.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{"user": "User"})
	// })

	// r.GET("/welcome", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "welcome.html", gin.H{"user": "User"})
	// })

	Rcontext.Run()
}
