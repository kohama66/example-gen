package repository

import (
	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/db/query"
)

type user struct{}

func NewUser() repository.User {
	return &user{}
}

func (u *user) GetByName(name string) (*entity.User, error) {
	us, err := query.User.Where(query.User.Username.Eq(name)).First()
	if err != nil {
		return nil, err
	}

	return entity.NewUser(int64(us.ID), us.Username, us.Email, us.CreatedAt), nil
}
