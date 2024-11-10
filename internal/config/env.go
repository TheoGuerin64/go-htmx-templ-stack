package config

import "fmt"

type Environnement struct {
	state string
}

const (
	environnementDev  = "dev"
	environnementProd = "prod"
)

func (e Environnement) IsDev() bool {
	return e.state == environnementDev
}

func (e Environnement) IsProd() bool {
	return e.state == environnementProd
}

func parseEnvironnement(value string) (interface{}, error) {
	var state string
	switch value {
	case environnementDev:
		state = environnementDev
	case environnementProd:
		state = environnementProd
	default:
		return "", fmt.Errorf("invalid environment value: %s", value)
	}
	return Environnement{state: state}, nil
}
