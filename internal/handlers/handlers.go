package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"MorzeText/internal/service"
)

func ParseHTML(w http.ResponseWriter, r *http.Request) {
	// парсим форму
	err := r.ParseMultipartForm(10 << 20) // mb ParseForm
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError) // mb better StatusBadRequest
		return
	}

	// извлекаем файл из формы
	myFile, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Form don't have a file", http.StatusInternalServerError) // mb better StatusBadRequest
		return
	}
	defer myFile.Close()

	// извлекаем данные из файла
	data, err := io.ReadAll(myFile)
	if err != nil {
		http.Error(w, "Reading error: can't get data from file", http.StatusInternalServerError)
		return
	}

	// конвертируем данные в строку
	strInput := string(data)

	// отправляем строку на конвертацию МорзеТекст
	str, err := service.Service(strInput)
	if err != nil {
		http.Error(w, "Error MorzeText convertation", http.StatusInternalServerError)
		return
	}

	// создать локальный файл и записать туда результаты конвертации
	fileName := time.Now().Format("15.04.05_02.01.2006")

	err = os.WriteFile(fileName, []byte(str), 0755)
	if err != nil {
		http.Error(w, "Failed to save result", http.StatusInternalServerError)
		return
	}
	fmt.Printf("File %s created\n", fileName)

	// получаем расширение файла (не понял зачем)
	// ext := filepath.Ext(fileName)

	// вернуть результат конвертации строки
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(str))
	if err != nil {
		http.Error(w, "Failed to write ResponseWriter", http.StatusInternalServerError)
		return
	}
}
