package user

type Service interface {
	GetCurrent(id string) (User, error)
	CreateNew(user User) (User, error)
	Login(email string, password string) (token string, err error)
	Update(id string, name string, address string) (User, error)
}

type Repository interface {
	CreateNew(user User) (User, error)
	FindById(email string) (User, error)
	FindByEmail(email string) (User, error)
	Update(id string, name string, address string) (User, error)
}
