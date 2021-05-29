package api

import (
	"hex/internal/ports"
)

// Dependency Injection
type APIAdapter struct {
	db ports.DbPort
	arith ports.ArithmeticPort
}

func NewApiAdapter(db ports.DbPort,arith ports.ArithmeticPort) *APIAdapter{
	return &APIAdapter{db: db ,arith: arith}
}

func (apia APIAdapter) GetAddition(a, b int32) (int32, error){
	response, err := apia.arith.Addition(a,b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(response, "addition")
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (apia APIAdapter) GetSubstraction(a, b int32) (int32, error){
	response, err := apia.arith.Substraction(a,b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(response, "substraction")
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (apia APIAdapter) GetMultiplication(a, b int32) (int32, error){
	response, err := apia.arith.Multiplication(a,b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(response, "multiplication")
	if err != nil {
		return 0, err
	}
	return response, nil
}

func (apia APIAdapter) GetDivision(a, b int32) (int32, error){
	response, err := apia.arith.Division(a,b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(response, "division")
	if err != nil {
		return 0, err
	}
	return response, nil
}

