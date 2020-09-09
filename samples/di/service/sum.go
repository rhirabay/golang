package service

type Sum struct {
}

func (c Sum) calculate(a int, b int) int {
	return a + b
}
