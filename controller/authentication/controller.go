package authentication

import "TruckMonitor-Backend/service"

type controller struct {
	authenticationService service.AuthenticationService
}

func Controller(authenticationService service.AuthenticationService) *controller {
	return &controller{
		authenticationService: authenticationService,
	}
}
