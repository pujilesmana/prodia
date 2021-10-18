package user

import (
	"errors"
	"log"

	"backend/business"
	"backend/business/utils"
)

type service struct {
	repository Repository
}

func NewUserService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetCurrent(id string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.InternalServerError)
	}

	return user, nil
}

func (s *service) CreateNew(user User) (User, error) {
	err := utils.GetValidator().Struct(user)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.BadRequest)
	}
	user.ID = utils.GenerateID()
	hashingPassword, err := utils.Hashing(user.Password)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.InternalServerError)
	}

	user.Password = hashingPassword
	userData, err := s.repository.CreateNew(user)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.BadRequest)
	}
	return userData, nil
}

func (s *service) Login(email string, password string) (string, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.Unauthorized)
	}

	err = utils.CompareHash(user.Password, password)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.Unauthorized)
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}
	return token, nil
}

func (s *service) Update(id, name, address string) (User, error) {
	user, err := s.repository.Update(id, name, address)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.InternalServerError)
	}

	return user, nil
}
