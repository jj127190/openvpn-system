package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func NoResponse(c *gin.Context) {
	// c.JSON(http.StatusNotFound, gin.H{
    //     "status": 404,
    //     "error":  "404 ,page not exists!",
    // })
	c.HTML(http.StatusNotFound, "error_return_index.html", gin.H{"msg":"请求地址错误!"})
}
