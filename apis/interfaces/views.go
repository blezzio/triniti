package interfaces

import "net/http"

type View interface {
	Exec(wr http.ResponseWriter, data any) error
	AddHeaders(wr http.ResponseWriter)
}
