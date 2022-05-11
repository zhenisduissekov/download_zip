package main

import (
	"github.com/zhenisduissekov/download_zip/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/downloadzip", handler.ZipHandler)
	http.ListenAndServe(":9099", nil)
}
