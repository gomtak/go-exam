package controller

type result int

func NewCTL() (result, error) {

	return result(0), nil
}

func (r result) CalcSum(a int, b int) int {
	return a + b
}
func (r result) CalcMul(a int, b int) int {

	return a * b
}
func (r result) CalcDiv(a int, b int) int {

	return a / b
}
func (r result) CalcSub(a int, b int) int {
	return a - b
}
