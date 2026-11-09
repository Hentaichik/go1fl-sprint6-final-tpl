package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	http.ServeFile(w, r, ".github/workflows/index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("failed to read file:%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := service.Translate(string(data))
	if len(res) == 0 {
		log.Println("func Translate return empty string")
		http.Error(w, "empty resault translation", http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(header.Filename)

	fileName := fmt.Sprintf("response_%s%s", time.Now().UTC().String(), ext)
	err = os.WriteFile(fileName, []byte(res), 0644)
	if err != nil {
		log.Printf("error writing file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}
