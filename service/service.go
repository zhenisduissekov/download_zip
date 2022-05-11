package service

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
)

func ZipFiles(w http.ResponseWriter, files []string) error {
	zw := zip.NewWriter(w)
	for _, val := range files {
		readNzipFile(zw, val)
	}
	// close the zip Writer to flush the contents to the ResponseWriter
	err := zw.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func readNzipFile(zw *zip.Writer, val string) error {
	f, err := os.Open(val)
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()

	// write straight to the http.ResponseWriter

	cf, err := zw.Create(f.Name())
	if err != nil {
		log.Println(err)
		return err
	}

	// copy the file contents to the zip Writer
	_, err = io.Copy(cf, f)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
