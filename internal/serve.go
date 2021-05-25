package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var PORT = 6137

type info struct {
	Port uint
	Name string
}

func Serve(c chan Api) int {
	//gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	v0 := engine.Group("/")
	var info info
	info.Port = uint(PORT)
	info.Name = "wisper"
	v0.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, `/Users/mardan/workspace/golang/wisper/web/out/index.html`, info)
	})
	vi := engine.Group("/api/v1")
	vi.GET("/refresh", func(ctx *gin.Context) {
		c <- Api{code: 1, op: "refresh", parms: ""}
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1",
		})
	})
	vi.POST("/update", func(ctx *gin.Context) {
		url := ctx.PostForm("url")
		c <- Api{code: 1, op: "update", parms: url}
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1",
			"url":  url,
		})
	})
	go engine.Run(":" + strconv.Itoa(PORT))

	return PORT
}
