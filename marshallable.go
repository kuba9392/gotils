package gotils

type marshallable interface {
	MarshalJSON() (string, error)
	UnmarshalJSON() error
}
