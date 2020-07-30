package service

// create login interface ... return bool

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	ValidUsername string
	ValidPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		ValidUsername: "Dinesh",
		ValidPassword: "Password",
	}
}

// implement method signature...

func (service *loginService) Login(username string, password string) bool {
	return service.ValidUsername == username &&
		service.ValidPassword == password
}
