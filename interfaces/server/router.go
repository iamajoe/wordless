package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iamajoe/wordless/entity"
	"github.com/iamajoe/wordless/httperr"
)

func accessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setEndpoints(r chi.Router, repos entity.Repositories) {
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		HandleResponse(w, true)
	})

	r.Route("/u", getUrlEndpoints(repos))
}

func GetParamID(r *http.Request, param string) (int, error) {
	idStr := chi.URLParam(r, param)
	if idStr == "" {
		return -1, httperr.NewError(
			http.StatusBadRequest,
			errors.New(fmt.Sprintf("%s required", param)),
		)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, httperr.NewError(http.StatusBadRequest, err)
	}

	return id, nil
}

func GetAPIRouter(r *chi.Mux, authSecret string, repos entity.Repositories) {
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.CleanPath)
		r.Use(middleware.NoCache)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(accessControl)

		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			HandleErrResponse(w, httperr.NewError(http.StatusMethodNotAllowed, errors.New("not allowed")))
		})
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			HandleErrResponse(w, httperr.NewError(http.StatusNotFound, errors.New("not allowed")))
		})

		setEndpoints(r, repos)
	})
}

func GetStaticRouter(r *chi.Mux) {
	root := "./app/dist"
	fs := http.FileServer(http.Dir(root))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

func GetRouter(authSecret string, repos entity.Repositories) *chi.Mux {
	r := chi.NewRouter()
	GetAPIRouter(r, authSecret, repos)
	GetStaticRouter(r)

	return r
}
