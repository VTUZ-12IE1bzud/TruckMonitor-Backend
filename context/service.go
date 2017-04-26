package context

import "TruckMonitor-Backend/service"

type (
	ServiceContext interface {
		AuthenticationService() service.AuthenticationService
		TokenService() service.TokenService
		UserService() service.UserService
	}

	serviceContext struct {
		authenticationService service.AuthenticationService
		tokenService          service.TokenService
		userService           service.UserService
	}
)

func NewServiceContext(tokenKey string, daoContext DaoContext) ServiceContext {
	userService := service.NewUserService(daoContext.EmployeeDao())
	tokenService := service.NewTokenService(tokenKey, daoContext.EmployeeDao())
	authenticationService := service.NewAuthenticationService(userService, tokenService)

	return &serviceContext{
		userService:           userService,
		tokenService:          tokenService,
		authenticationService: authenticationService,
	}

}

func (s serviceContext) AuthenticationService() service.AuthenticationService {
	return s.authenticationService
}

func (s serviceContext) TokenService() service.TokenService {
	return s.tokenService
}

func (s serviceContext) UserService() service.UserService {
	return s.userService
}
