package main

import (
	"fmt"
	"math"
)

func main() {
	//var investmentAmount = 1000
	//var expectedReturn = 5.5
	//var years = 10
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5
	// var years float64 = 10

	fmt.Scan(&investmentAmount)

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value is ", futureValue)
	fmt.Println(futureRealValue)
}
