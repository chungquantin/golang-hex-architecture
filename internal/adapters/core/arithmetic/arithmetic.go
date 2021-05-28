package arithmetic

type ArithmeticAdapter struct {
}

func NewArithmeticAdapter() *ArithmeticAdapter {
	return &ArithmeticAdapter{}
}

func (arith ArithmeticAdapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith ArithmeticAdapter) Substraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith ArithmeticAdapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith ArithmeticAdapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
