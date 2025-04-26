package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// создал структуру сервера
type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

// создаёт новый экземпляр структуры http.Server
func New(logger *log.Logger) *Server {
	// Создаем роутер
	router := http.NewServeMux()

	// зарег хендлеры
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Обработка запроса на /")
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Главная страница конвертер Морзе-Текст"))
	})

	router.HandleFunc("/convert", handlers.ParseHTML)

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
		httpServer: httpServer,
	}
}
