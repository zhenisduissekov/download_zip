package service

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func ZipFiles(c *gin.Context, files []string) (int, error) {
	count := len(files)
	zw := zip.NewWriter(c.Writer)
	for _, val := range files {
		err := readNzipFile(zw, val)
		if err != nil {
			count -= 1
		}
	}

	err := zw.Close()
	if err != nil {
		log.Println(err)
		return count, err
	}
	return count, nil
}

func readNzipFile(zw *zip.Writer, val string) error {
	f, err := os.Open(val)
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()

	cf, err := zw.Create(f.Name())
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = io.Copy(cf, f)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
