package Service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type LoginInformation struct {
	email string
	password string
}

func StaticLoginService() LoginService  {
	return &LoginInformation{
		email:    "hakim@nganu.com",
		password: "passwordyangsusah",
	}
}

func (l *LoginInformation)LoginUser(email string, password string)bool  {
	return l.email == email && l.password == password
}
