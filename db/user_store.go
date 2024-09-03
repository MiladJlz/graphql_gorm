package db

import (
	"context"
	"fmt"

	"github.com/miladjlz/golang-graphql-gorm-postgresql/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUserByID(int) (*model.User, error)
	GetUsers() ([]*model.User, error)
	InsertUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(int) (*model.DeleteUserResponse, error)
	UpdateUser(int, model.UpdateUserInput) (*model.User, error)
}

type PostgresUserStore struct {
	client *gorm.DB
}

func NewPostgresUserStore() (*PostgresUserStore, error) {
	config := GetPostgresConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return &PostgresUserStore{db}, nil
}

func (s *PostgresUserStore) GetUser(id int) (*model.User, error) {
	user := model.User{}
	if err := s.client.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (s *PostgresUserStore) GetUsers() ([]*model.User, error) {
	users := []*model.User{}

	if err := s.client.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}

func (s *PostgresUserStore) InsertUser(input model.CreateUserInput) (*model.User, error) {
	user := model.User{

		Name:  input.Name,
		Email: input.Email,
	}

	result := s.client.Table("users").Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (s *PostgresUserStore) UpdateUser(id int, input model.UpdateUserInput) (*model.User, error) {
	var user model.User
	if err := s.client.Table("users").First(&user, id).Error; err != nil {
		return nil, err
	}

	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Name != nil {
		user.Name = *input.Name
	}

	if err := s.client.Save(user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (s *PostgresUserStore) DeleteUser(id int) (*model.DeleteUserResponse, error) {
	var user model.User

	if err := s.client.First(&user, id).Delete(id).Error; err != nil {
		return nil, err
	}

	return &model.DeleteUserResponse{DeleteUserID: id}, nil

}
