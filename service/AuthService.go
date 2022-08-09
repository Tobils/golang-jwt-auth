package service

type AuthService interface {
	Login(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticAuthService() AuthService {
	return &loginInformation{
		email:    "suhada@gmail.com",
		password: "testing",
	}
}

func (info *loginInformation) Login(email string, password string) bool {
	return info.email == email && info.password == password
}
