package ports

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubstraction()
	GetMultiplication()
	GetDivision()
}