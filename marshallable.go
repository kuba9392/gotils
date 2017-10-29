package gotils

type marshallable interface {
	jsonSerialize() map[string]interface{}
}
