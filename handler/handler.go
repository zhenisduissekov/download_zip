package handler

import (
	"encoding/json"
	"github.com/zhenisduissekov/download_zip/service"
	"io/ioutil"
	"log"
	"net/http"
)

func ZipHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Println("wrong method used")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("wrong method used")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error at ioutil.ReadAll", err)
		return
	}
	var files struct {
		Body []string `json:"body"`
	}

	err = json.Unmarshal(body, &files)
	if err != nil {
		log.Println("error Unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = service.ZipFiles(w, files.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "files.zip")

	_ = json.NewEncoder(w).Encode("finished")
}
