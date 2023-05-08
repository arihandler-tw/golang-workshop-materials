package mem_store

import "fmt"

type Store[K comparable, V any] interface {
	// Find takes an ID of type K.
	// If there's an element with that ID, Find returns a pointer to that element and `true`.
	// If there is no element with that ID, Find returns `nil`, `false`.
	Find(id K) (*V, bool)

	// Store takes an ID and a value to store under that ID.
	// If successful, it returns a pointer to the stored value and `nil`.
	// If it's unable to store the value, it returns `nil` and a non-nil error.
	// If the ID is already present in the store, it's value will be overwritten.
	Store(id K, value V) (*V, error)
}

type MemStore[K comparable, V any] struct {
	store map[K]V
}

func NewMemStore[K comparable, V any]() Store[K, V] {
	return &MemStore[K, V]{store: make(map[K]V)}
}

func (s *MemStore[K, V]) Find(id K) (*V, bool) {
	if value, ok := s.store[id]; ok {
		return &value, true
	}
	return nil, false
}

func (s *MemStore[K, V]) Store(id K, value V) (*V, error) {
	s.store[id] = value
	return &value, nil
}

// MemStoreWithValidation a MemStore that will only store a value if it passes the given `validate()` function
type MemStoreWithValidation[K comparable, V any] struct {
	*MemStore[K, V]
	validate func(V) error
}

func NewMemStoreWithValidation[K comparable, V any](validator func(val V) error) Store[K, V] {
	return &MemStoreWithValidation[K, V]{
		&MemStore[K, V]{store: make(map[K]V)},
		validator,
	}
}

func (s *MemStoreWithValidation[K, V]) Find(id K) (*V, bool) {
	return s.MemStore.Find(id)
}

func (s *MemStoreWithValidation[K, V]) Store(id K, value V) (*V, error) {
	if err := s.validate(value); err != nil {
		return nil, fmt.Errorf("value is invalid: %w", err)
	}

	return s.MemStore.Store(id, value)
}
