package go_tils

type marshallable interface {
	MarshalJSON() (string, error)
	UnmarshalJSON() error
}
