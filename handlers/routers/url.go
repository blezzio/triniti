package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/blezzio/triniti/handlers/interfaces"
	"github.com/blezzio/triniti/handlers/types"
	"github.com/blezzio/triniti/utils"
)

type URL struct {
	service interfaces.URLUseCase
	index   interfaces.View
}

func NewURL(service interfaces.URLUseCase, index interfaces.View) *URL {
	return &URL{service: service, index: index}
}

func (h *URL) Route(serveMux *http.ServeMux) {
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

	resp := []byte(fmt.Sprintf("<h1>localhost:4444/%v<h1>", hash))
	if _, err := w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	if err := h.index.Exec(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.index.AddHeaders(w)
}
