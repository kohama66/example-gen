package repository

import (
	"api/domain/entity"
	"context"
)

type User interface {
	GetByName(name string) (*entity.User, error)
	AddCreditCard(ctx context.Context, userID int64, cardNumber string) error
}
