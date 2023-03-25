package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/iamajoe/wordless/config"
	"github.com/iamajoe/wordless/domain"
	"github.com/iamajoe/wordless/entity"
)

func reqCreateUrl(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var urlRaw struct {
			Url  string `json:"url"`
			Mini bool   `json:"mini"`
		}
		err := json.NewDecoder(r.Body).Decode(&urlRaw)
		if err != nil {
			HandleErrResponse(w, err)
			return
		}

		id, err := domain.CreateUrl(urlRaw.Url, urlRaw.Mini, repos.GetUrl())
		if err != nil {
			HandleErrResponse(w, err)
			return
		}

		HandleResponse(w, id)
	}
}

// TODO: need to test the fetch
func reqFetchId(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if len(id) == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		url, err := domain.FetchId(id, repos.GetUrl())
		if err != nil || len(url) == 0 {
			// hide the error from the user
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		c, err := config.Get(os.Getenv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// DEV: do not redirect when testing
		if c.Env == config.EnvTesting {
			HandleResponse(w, url)
			return
		}

		// time to redirect to the right place
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func getUrlEndpoints(repos entity.Repositories) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", reqCreateUrl(repos))
		r.Get("/{id}", reqFetchId(repos))
	}
}
