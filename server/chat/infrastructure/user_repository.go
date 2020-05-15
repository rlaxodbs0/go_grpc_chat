package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DBName = "grpc_chat"
)

type userRepository struct {
	db *gorm.DB
}

var _ repository.UserRepositoryImpl = &userRepository{} //interface impl

func NewUserRepository() repository.UserRepositoryImpl {
	db, err := gorm.Open("sqlite3", "user_db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	AutoMigrate(db)
	return &userRepository{db:db}
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}

func (ur *userRepository) CreateUser(user *model.User) error {
	tx := ur.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	if err := tx.Where("username=?",user.Username).Find(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (ur *userRepository) UpdateUser(user *model.User) error {
	tx := ur.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Model(user).Update(user).Error; err != nil {
		return err
	}
	if err := tx.Where("username=?", user.Username).Find(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (ur *userRepository) DeleteUser(user *model.User) error {
	return ur.db.Delete(user).Error
}

func (ur *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var m model.User
	err := ur.db.Where("username=?", username).Preload("Friends").Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (ur *userRepository) AddFriend(user *model.User, friend *model.User) error{
	tx := ur.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Model(user).Association("Friends").Find(friend).Error; err != nil {
		return err
	}
	if err := tx.Model(user).Association("Friends").Append(friend).Error; err != nil {
		return err
	}
	if err := tx.Where("username=?", friend.Username).Preload("User").First(friend).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (ur *userRepository) RemoveFriend(user *model.User, friend *model.User) error{
	return ur.db.Model(user).Association("Friends").Delete(friend).Error
}
