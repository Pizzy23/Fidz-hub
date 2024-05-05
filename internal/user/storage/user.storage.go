package storage

import (
	"fanify/db"
	inter "fanify/internal/user/interface"
	"time"

	"gorm.io/gorm"
)

func ConsultUser(q *gorm.DB, email string) (*db.User, error) {
	var user db.User
	if err := q.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(q *gorm.DB, input inter.UserControllerInputDb) (*db.User, error) {
	newUser := db.User{
		Email:     input.Email,
		Password:  input.Password,
		Wallet:    input.Wallet,
		WalletId:  input.WalletId,
		ProjectId: input.ProjectId,
		CreateAt:  time.Now(),
	}
	if err := q.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}
