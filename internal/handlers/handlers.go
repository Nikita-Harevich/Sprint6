package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HTML(ww http.ResponseWriter, rr *http.Request) {
	http.ServeFile(ww, rr, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10000000000000); err != nil {
		log.Printf("Ошибка парсинга: %v", err)
		http.Error(w, "ошибка парсинга", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка получения формы из файла: %v", err)
		http.Error(w, "ошибка получения формы", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	resultText, err := service.ConvertSrting(string(fileBytes))
	if err != nil {
		log.Printf("Ошибка resultText1")
		http.Error(w, "Ошибка resultTest2", http.StatusInternalServerError)
		return
	}

	timeFile := time.Now().UTC().Format("20060102150405")
	fileExt := filepath.Ext(header.Filename)
	newFile := timeFile + fileExt

	err = os.WriteFile(newFile, []byte(resultText), 0755)
	if err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
	}

	w.Write([]byte(resultText))
}
