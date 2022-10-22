package presentation

import (
	"chi_sample/config"
	"chi_sample/presentation/controller"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CreateRoute() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Enviroment.AllowOrigin},
		AllowedMethods:   []string{"POST"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Timeout(time.Second * 10))
	r.Use(middleware.Recoverer)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("存在しないエンドポイントです。"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("許可されていないメソッドです。"))
	})

	r.Group(func(r chi.Router) {
		ac := controller.CreateAccountController()
		r.Mount("/account", ac)
	})

	return r
}
