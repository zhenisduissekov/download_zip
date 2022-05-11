package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenisduissekov/download_zip/handler"
)

var (
	port = ":9099"
)

func main() {

	r := gin.Default()
	r.POST("/downloadzip", handler.ZipHandler)
	r.Run(port)
}
