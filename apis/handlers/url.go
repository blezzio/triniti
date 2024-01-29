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
	FailureView ViewName = "FAILURE"
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
		h.showErrorPage(w, req, http.StatusMethodNotAllowed, fmt.Errorf("%v method is not allowed", req.Method))
		return
	}
}

func (h *URL) handlePost(w http.ResponseWriter, req *http.Request) {
	formURL := req.FormValue("url")
	_, err := url.ParseRequestURI(formURL)
	if err != nil {
		h.showErrorPage(w, req, http.StatusBadRequest, err)
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
		uri = h.fix(uri)
		h.getHash(w, req, uri)
	}
}

func (h *URL) getHash(w http.ResponseWriter, req *http.Request, url string) {
	hash, err := h.service.GetHash(req.Context(), url)
	if err != nil {
		if err != nil {
			h.showErrorPage(w, req, http.StatusInternalServerError, err)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)

	shortenedURL := fmt.Sprintf("%s/%v", strings.ToLower(os.Getenv("TRINITI_URL")), hash)

	successView, ok := h.views[SuccessView]
	if !ok {
		if _, err := w.Write([]byte(shortenedURL)); err != nil {
			h.showErrorPage(w, req, http.StatusInternalServerError, fmt.Errorf("no %s view", SuccessView))
		}
		return
	}
	if err := successView.Exec(w, &types.HTMLSuccessView{
		AcceptLanguage: req.Header.Get("Accept-Language"),
		URL:            shortenedURL,
	}); err != nil {
		h.showErrorPage(w, req, http.StatusInternalServerError, err)
		return
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
			h.showErrorPage(w, req, http.StatusNotFound, err)
			return
		}
		h.showErrorPage(w, req, http.StatusInternalServerError, err)
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
		h.showErrorPage(w, req, http.StatusInternalServerError, fmt.Errorf("no %s view", IndexView))
		return
	}

	if err := index.Exec(w, data); err != nil {
		h.showErrorPage(w, req, http.StatusInternalServerError, err)
		return
	}
}

func (h *URL) showErrorPage(w http.ResponseWriter, req *http.Request, code int, err error) {
	data := &types.HTMLErrorView{
		AcceptLanguage: req.Header.Get("Accept-Language"),
		Code:           code,
		Error:          err,
	}

	errView, ok := h.views[FailureView]
	if !ok {
		http.Error(
			w,
			utils.Trace(
				fmt.Errorf("no %s view", FailureView), "also previous error %w", err,
			).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	if err := errView.Exec(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
