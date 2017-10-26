package go_tils

import "net/http"

type wrapper interface {
	Wrap(w *http.ResponseWriter)
}