package main

import "golang/samples/di/service"

func main() {
	sum := service.Sum{}
	service := service.CalculateService{Calculator: sum}

	service.Execute()
}
