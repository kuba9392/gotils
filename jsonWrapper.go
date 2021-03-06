//Package gotils contains golang utilities to make designing new services more cleaner and easier.
package gotils

import (
	"encoding/json"
	"github.com/stretchr/codecs/constants"
	"net/http"
)

//jsonWrapper is a type which helps you with many processes related with JSON format
//For example, you can wrap HTTP response to JSON content-type or quickly encode self-defined type.
type jsonWrapper struct{}

//Constructor for jsonWrapper
func NewJsonWrapper() *jsonWrapper {
	wrapper := jsonWrapper{}
	return &wrapper
}

//This function allows you to wrap created response to JSON format
func (w *jsonWrapper) Wrap(rw http.ResponseWriter) *jsonWrapper {
	rw.Header().Set("Content-Type", constants.ContentTypeJSON)
	return w
}

//Encode creates JSON message from self-defined type which contains Marshallable interface
func (w *jsonWrapper) Encode(m Marshallable) (string, error) {
	serialized := m.JsonSerialize()
	encoded, err := json.Marshal(serialized)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

//Decode returns map with data from JSON message
//You can implement this function to createFromArray method in your Marshallable type
func (w *jsonWrapper) Decode(message string) (map[string]interface{}, error) {
	decoded := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &decoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
