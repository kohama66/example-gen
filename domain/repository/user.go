package repository

import "api/domain/entity"

type User interface {
	GetByName(name string) (*entity.User, error)
}
