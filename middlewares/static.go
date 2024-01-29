package middlewares

import (
	"embed"
	"mime"
	"net/http"
	"strings"
)

type Static struct {
	handler http.Handler
	fs      embed.FS
	path    string
}

func NewStatic(fs embed.FS, path string) *Static {
	return &Static{fs: fs, path: strings.Trim(path, "/")}
}

func (mw *Static) Wrap(h http.Handler) http.Handler {
	return &Static{
		handler: h,
		fs:      mw.fs,
		path:    mw.path,
	}
}

func (mw *Static) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path[0] == '/' {
		path = path[1:]
	}
	if len(path) == 0 {
		mw.handler.ServeHTTP(rw, req)
		return
	}

	parts := strings.Split(path, "/")
	if len(parts) < 2 || parts[0] != mw.path {
		mw.handler.ServeHTTP(rw, req)
		return
	}

	mw.handle(rw, path, parts)
}

func (mw *Static) handle(rw http.ResponseWriter, path string, parts []string) {
	if b, err := mw.fs.ReadFile(path); err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	} else {
		last := parts[len(parts)-1]
		ext := strings.Split(last, ".")
		t := mime.TypeByExtension("." + ext[1])
		rw.Header().Add("Content-Type", t)
		if _, err = rw.Write(b); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
