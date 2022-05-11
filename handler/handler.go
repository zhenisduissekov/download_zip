package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zhenisduissekov/download_zip/service"
	"log"
	"net/http"
)

var request struct {
	Body []string `json:"body"`
}

var response struct {
	Message string `json:"message"`
}

func ZipHandler(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Println("error at ShouldBindBodyWith", err)
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println("request.Body", request.Body)

	if len(request.Body) == 0 {
		response.Message = "no files to zip"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := service.ZipFiles(c, request.Body);
	if err != nil {
		log.Println("error at zipfiles", err)
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.Writer.Header().Set("Content-type", "application/zip")
	c.Writer.Header().Set("Content-Disposition", "files.zip")

	response.Message = fmt.Sprintf("%d files zipped with success", count)
	c.JSON(http.StatusOK, response)
}
