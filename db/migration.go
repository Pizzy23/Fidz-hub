package db

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	Wallet    string    `gorm:"column:wallet" json:"wallet"`
	WalletId  string    `gorm:"column:walletId" json:"walletId"`
	ProjectId string    `gorm:"column:projectId" json:"projectId"`
	CreateAt  time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;not null" json:"update_at"`
}
