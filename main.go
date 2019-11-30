// Package main ...
// Date: 2019/05/02
// Author: gtlions.l@qq.com
package main

import (
	"gokv/model"
	"gokv/router"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	e := router.InitRouter()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{Data: map[string]string{"Service": "Pong"}})
	})
	e.Run("127.0.0.1:8080")
}
