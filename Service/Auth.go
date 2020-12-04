package Service

type AuthInformation struct {
	Email string
	Permission []string
	Role string
}

func StaticAuthService() []AuthInformation{
	var authlist = []AuthInformation{
		AuthInformation{
			Role:      "admin",
			Permission: []string{"/admin","/user"},
		},
		AuthInformation{
			Role:      "user",
			Permission: []string{"/user"},
		},
		AuthInformation{
			Role:      "waiter",
			Permission: []string{"/waiter"},
		},
	}
	return authlist


}