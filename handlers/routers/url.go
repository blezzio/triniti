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

func (h *URL) handle(w http.ResponseWriter, res *http.Request) {
	uri := res.RequestURI
	if uri[0] == '/' {
		uri = uri[1:]
	}
	if len(uri) == 0 {
		h.showIndexPage(w, res)
		return
	}

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		h.getFullURL(w, res, uri)
	} else {
		h.getHash(w, res, uri)
	}
}

func (h *URL) getHash(w http.ResponseWriter, res *http.Request, url string) {
	url = h.fix(url)
	hash, err := h.service.GetHash(res.Context(), url)
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

func (h *URL) getFullURL(w http.ResponseWriter, res *http.Request, hash string) {
	url, err := h.service.GetFullURL(res.Context(), hash)
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

func (h *URL) showIndexPage(w http.ResponseWriter, res *http.Request) {
	data := &types.HTMLIndexView{
		AcceptLanguage: res.Header.Get("Accept-Language"),
	}

	if err := h.index.Exec(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.index.AddHeaders(w)
}
