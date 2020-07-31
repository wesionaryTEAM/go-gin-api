package service

// create login interface ... return boolpackage service

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	ValidEmail    string
	ValidPassword string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		ValidEmail:    "dinesh.silwal@wesionary.team",
		ValidPassword: "nothing",
	}
}
func (service *loginInformation) LoginUser(email string, password string) bool {
	return service.ValidEmail == email && service.ValidPassword == password
}
