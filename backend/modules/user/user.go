package user

import (
	"errors"
	"time"

	userBusiness "backend/business/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateNew(user userBusiness.User) (userBusiness.User, error) {
	err := repo.db.Where("email = ?", user.Email).First(&User{}).Error
	if err == nil {
		return userBusiness.User{}, errors.New("email already exist")
	}

	userData := convertToUserModel(user)

	err = repo.db.Create(&userData).Error
	if err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(userData)

	return userBusiness, nil
}

func (repo *UserRepository) FindById(userId string) (userBusiness.User, error) {
	var userData User
	err := repo.db.Where("id = ?", userId).First(&userData).Error
	if err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(userData)
	return userBusiness, nil
}

func (repo *UserRepository) FindByEmail(email string) (userBusiness.User, error) {
	var user User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func (repo *UserRepository) Update(userId, name, address string) (userBusiness.User, error) {
	var user User
	err := repo.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return userBusiness.User{}, err
	}

	err = repo.db.Model(&user).Updates(&User{Name: name, Address: address}).Error
	if err != nil {
		return userBusiness.User{}, err
	}

	userBusiness := convertToUserBusiness(user)
	return userBusiness, nil
}

func convertToUserModel(user userBusiness.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func convertToUserBusiness(user User) userBusiness.User {
	return userBusiness.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
