package gotils

type Marshallable interface {
	JsonSerialize() map[string]interface{}
}
