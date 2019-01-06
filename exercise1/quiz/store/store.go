package store

// Store interface that models a store
type Store interface {
	Get() ([]byte, error)
	Set([]byte) error
}
