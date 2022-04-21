package main

import (
	"github.com/jktan0504/bmi_func/info"
)

func main() {
	info.PrintGreetingMessages()
	weight, height := getUserMetrics()
	info.PrintResult(weight, height, bmiCaculator(weight, height))
}

func bmiCaculator(weight float64, height float64) (bmi float64) {
	bmi = weight / (height * height)
	return
}
