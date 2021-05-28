package api

import (
	"hex/internal/ports"
)

// Dependency Injection
type APIAdapter struct {
	arith ports.ArithmeticPort
}

func NewApiAdapter() *APIAdapter{
	return &APIAdapter{}
}

func (api APIAdapter) GetAddition(a, b int32) (int32, error){
	response, err := api.arith.Addition(a,b)
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (api APIAdapter) GetSubstraction(a, b int32) (int32, error){
	response, err := api.arith.Substraction(a,b)
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (api APIAdapter) GetMultiplication(a, b int32) (int32, error){
	response, err := api.arith.Multiplication(a,b)
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (api APIAdapter) GetDivision(a, b int32) (int32, error){
	response, err := api.arith.Division(a,b)
	if err != nil {
		return 0, err
	}
	return response, nil
}

