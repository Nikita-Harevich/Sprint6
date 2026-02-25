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

func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("../index.html")
	http.ServeFile(w, r, filepath)
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Ошибка парсинга: %v", err)
		http.Error(w, "ошибка парсинга", http.StatusBadRequest)
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

	resultFile, err := os.Create(newFile)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
		return
	}

	defer resultFile.Close()

	file, err = os.OpenFile(newFile, os.O_WRONLY, 0755)
	if err != nil {
		log.Printf("Ошибка записи данных в файл: %v", err)
		http.Error(w, "ошибка записи в файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Write([]byte(resultText))
}
