package service

import "fmt"

type CalculateService struct {
	Calculator Calculator
}

func (service CalculateService) Execute() {
	var a int
	var b int
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)

	fmt.Println(service.Calculator.calculate(a, b))
}
