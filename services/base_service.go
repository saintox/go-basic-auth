package services

type Service struct {
	Login LoginService
	//Register RegisterService
}

func NewService() *Service {
	return &Service{
		Login: NewLoginService(),
	}
}
