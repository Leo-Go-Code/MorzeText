package server

import (
	"log"
	"net/http"
	"time"

	"MorzeText/internal/handlers"
)

// создал структуру сервера
type Server struct {
	logger     *log.Logger
	HTTPServer *http.Server
}

// создаёт новый экземпляр структуры http.Server
func NewServer(logger *log.Logger) *Server {
	// Создаем роутер
	router := http.NewServeMux()

	// зарег хендлеры
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/upload", handlers.ParseHTML) // проверяю метод в хендлере

	// Создаем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// возвращаем ссылку на сервер
	return &Server{
		logger:     logger,
		HTTPServer: httpServer,
	}
}
