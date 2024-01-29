package middlewares

import (
	"embed"
	"mime"
	"net/http"
	"strings"
)

type FavIco struct {
	handler http.Handler
	fs      embed.FS
	root    string
}

func NewFavIco(fs embed.FS, rootFolder string) *FavIco {
	return &FavIco{fs: fs, root: strings.Trim(rootFolder, "/")}
}

func (mw *FavIco) Wrap(h http.Handler) http.Handler {
	return &FavIco{
		handler: h,
		fs:      mw.fs,
		root:    mw.root,
	}
}

func (mw *FavIco) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/favicon.ico" && req.URL.Path != "/site.webmanifest" {
		mw.handler.ServeHTTP(rw, req)
		return
	}
	mw.handle(rw, req.URL.Path)
}

func (mw *FavIco) handle(rw http.ResponseWriter, path string) {
	if b, err := mw.fs.ReadFile(mw.root + path); err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	} else {
		if _, err = rw.Write(b); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
	ext := strings.Split(path, ".")
	t := mime.TypeByExtension("." + ext[1])
	rw.Header().Add("Content-Type", t)
}
