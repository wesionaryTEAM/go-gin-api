package service

// create login interface ... return boolpackage service

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "dinesh.silwal@wesionary.team",
		password: "nothing",
	}
}
func (service *loginInformation) LoginUser(email string, password string) bool {
	return service.email == email && service.password == password
}
