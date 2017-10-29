package gotils

import (
	"net/http"
	"github.com/stretchr/codecs/constants"
	"encoding/json"
)

type jsonWrapper struct {}

func NewJsonWrapper() *jsonWrapper {
	return &jsonWrapper{}
}

func (w *jsonWrapper) Wrap(rw http.ResponseWriter) *jsonWrapper {
	rw.Header().Set("Content-Type", constants.ContentTypeJSON)
	return w
}

func (w *jsonWrapper) Encode(m marshallable) (string, error) {
	serialized := m.jsonSerialize()
	encoded, err := json.Marshal(serialized)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (w *jsonWrapper) Decode(message string) (map[string]interface{}, error) {
	decoded := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &decoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}