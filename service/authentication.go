package service

import "TruckMonitor-Backend/model"

type AuthenticationService interface {
	Authenticate(email string, password string) (*model.Employee, error)
	CreateToken(email string, password string) (string, error)
	ResolveToken(sessionToken string) (*model.Employee, error)
}

type authenticationService struct {
	userService  UserService
	tokenService TokenService
}

func NewAuthenticationService(userService UserService, tokenService TokenService) AuthenticationService {
	return &authenticationService{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (s authenticationService) Authenticate(email string, password string) (*model.Employee, error) {
	user, err := s.userService.Validate(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s authenticationService) CreateToken(email string, password string) (string, error) {
	user, err := s.userService.Validate(email, password)
	if err != nil {
		return "", err
	}
	return s.tokenService.Create(user)
}

func (s authenticationService) ResolveToken(sessionToken string) (*model.Employee, error) {
	token, err := s.tokenService.Resolve(sessionToken)
	if err != nil {
		return nil, err
	}
	return s.userService.Get(token.EmployeeId)
}
