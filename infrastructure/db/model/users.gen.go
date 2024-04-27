// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         int32         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string        `gorm:"column:username;not null" json:"username"`
	Email      string        `gorm:"column:email;not null" json:"email"`
	CreatedAt  time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreditCard []*CreditCard `json:"credit_card"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
