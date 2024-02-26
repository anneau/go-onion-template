package model

import "github.com/gofrs/uuid"

type ID[T any] struct {
	Value uuid.UUID
}

func NewID[T any]() (ID[T], error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return ID[T]{}, err
	}

	return ID[T]{Value: uid}, nil
}

func NewIDFrom[T any](v string) (ID[T], error) {
	uid, err := uuid.FromString(v)
	if err != nil {
		return ID[T]{}, err
	}

	return ID[T]{Value: uid}, nil
}
