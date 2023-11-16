package repository

import (
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (t *userRepository) GetAUser(user *entities.User, id string) (err error) {
	if err := t.conn.Raw("SELECT * FROM users WHERE id = ?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (t *userRepository) GetAUserByFirebaseId(user *entities.User, firebaseId string) (err error) {
	if err := t.conn.Raw("SELECT * FROM users WHERE firebase_id = ?", firebaseId).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (t *userRepository) UpdateAUser(user *entities.User) (err error) {
	if err := t.conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
