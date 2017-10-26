package go_tils

import (
	"net/http"
	"github.com/stretchr/codecs/constants"
)

type jsonWrapper struct {}

func NewJsonWrapper() *jsonWrapper {
	return &jsonWrapper{}
}

func (w *jsonWrapper) Wrap(rw http.ResponseWriter) *jsonWrapper {
	rw.Header().Set("Content-Type", constants.ContentTypeJSON)
	return w
}

func (w *jsonWrapper) Encode(marshallable marshallable) (string, error) {
	return "", nil
}