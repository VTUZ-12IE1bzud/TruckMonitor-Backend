package common

type commonController struct{}

func Controller() *commonController {
	return &commonController{}
}
