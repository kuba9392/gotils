package gotils

import "net/http"

type wrapper interface {
	Wrap(w *http.ResponseWriter)
}