package gotils

type Marshallable interface {
	jsonSerialize() map[string]interface{}
}
