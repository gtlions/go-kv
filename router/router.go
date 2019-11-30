// Package router ...
// Date: 2019/05/02
// Author: gtlions.l@qq.com
package router

import (
	"gokv/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cache model.Cache

// InitRouter initial HTTP Engine
func InitRouter() *gin.Engine {
	cache.Data = make(map[string]string)
	e := gin.Default()
	e.Any("/", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			var key string
			p := c.Request.URL.Query()
			if len(p) != 1 {
				c.JSON(http.StatusBadRequest, model.Response{Code: 10400, Error: "Invalid Request."})
				c.Abort()
				return
			}
			for k := range p {
				key = k
			}
			if v, ok := cache.Data[key]; ok {
				c.JSON(http.StatusOK, model.Response{Data: v})
			} else {
				c.JSON(http.StatusBadRequest, model.Response{Code: 10400, Error: "Key Not Found."})
			}
			break
		case "POST":
			var d model.KV
			if err := c.ShouldBind(&d); err != nil {
				c.JSON(http.StatusBadRequest, model.Response{Code: 10400, Error: "Invalid Request."})
				c.Abort()
				return
			}
			cache.Data[d.Key] = d.Value
			break
		case "DELETE":
			var key string
			p := c.Request.URL.Query()
			if len(p) != 1 {
				c.JSON(http.StatusBadRequest, model.Response{Code: 10400, Error: "Invalid Request."})
				c.Abort()
				return
			}
			for k := range p {
				key = k
			}
			if _, ok := cache.Data[key]; ok {
				delete(cache.Data, key)
				c.JSON(http.StatusOK, model.Response{})
			} else {
				c.JSON(http.StatusBadRequest, model.Response{Code: 10400, Error: "Key Not Found."})
			}
			break
		}
	})
	return e
}
