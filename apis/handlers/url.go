package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/blezzio/triniti/apis/interfaces"
	"github.com/blezzio/triniti/apis/types"
	"github.com/blezzio/triniti/utils"
)

type ViewName string

const (
	IndexView   ViewName = "INDEX"
	SuccessView ViewName = "SUCCESS"
	ErrorView   ViewName = "ERROR"
)

type URL struct {
	service interfaces.URLUseCase
	views   map[ViewName]interfaces.View
}

func NewURL(
	service interfaces.URLUseCase, opts ...UrlOpt,
) *URL {
	u := &URL{service: service, views: map[ViewName]interfaces.View{}}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

func (h *URL) Build(serveMux *http.ServeMux) {
	serveMux.HandleFunc("/", h.handle)
}

func (h *URL) handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.handleGet(w, req)
	case http.MethodPost:
		h.handlePost(w, req)
	default:
		w.Header().Add("Allow", fmt.Sprintf("%v, %v", http.MethodGet, http.MethodPost))
		http.Error(w, fmt.Sprintf("%v method is not allowed", req.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (h *URL) handlePost(w http.ResponseWriter, req *http.Request) {
	formURL := req.FormValue("url")
	_, err := url.ParseRequestURI(formURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		h.getHash(w, req, formURL)
	}
}

func (h *URL) handleGet(w http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	if uri[0] == '/' {
		uri = uri[1:]
	}
	if len(uri) == 0 {
		h.showIndexPage(w, req)
		return
	}

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		h.getFullURL(w, req, uri)
	} else {
		h.getHash(w, req, uri)
	}
}

func (h *URL) getHash(w http.ResponseWriter, req *http.Request, url string) {
	url = h.fix(url)
	hash, err := h.service.GetHash(req.Context(), url)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)

	shortenedURL := fmt.Sprintf("%s/%v", strings.ToLower(os.Getenv("TRINITI_URL")), hash)

	successView, ok := h.views[SuccessView]
	if !ok {
		if _, err := w.Write([]byte(shortenedURL)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := successView.Exec(w, &types.HTMLSuccessView{
		AcceptLanguage: req.Header.Get("Accept-Language"),
		URL:            shortenedURL,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successView.AddHeaders(w)
}

func (h *URL) fix(url string) string {
	return strings.Replace(
		url, ":/", "://", 1,
	)
}

func (h *URL) getFullURL(w http.ResponseWriter, req *http.Request, hash string) {
	url, err := h.service.GetFullURL(req.Context(), hash)
	if err != nil {
		if terr, ok := err.(utils.TraceError); ok && terr.Is(sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusMovedPermanently)
}

func (h *URL) showIndexPage(w http.ResponseWriter, req *http.Request) {
	data := &types.HTMLIndexView{
		AcceptLanguage: req.Header.Get("Accept-Language"),
	}

	index, ok := h.views[IndexView]
	if !ok {
		http.Error(w, "no INDEX view", http.StatusInternalServerError)
		return
	}

	if err := index.Exec(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	index.AddHeaders(w)
}
