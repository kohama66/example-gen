// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCreditCard = "CreditCard"

// CreditCard mapped from table <CreditCard>
type CreditCard struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int32  `gorm:"column:user_id;not null" json:"user_id"`
	CardNumber string `gorm:"column:card_number;not null" json:"card_number"`
}

// TableName CreditCard's table name
func (*CreditCard) TableName() string {
	return TableNameCreditCard
}
