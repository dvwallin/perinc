package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	err               error
	percentage        float64
	initialPercentage float64
	period            int64
	inputNumber       string
	inputPercentage   string
	inputPeriod       string
	result            float64
	i                 int64
)

func main() {

	if len(os.Args) < 4 {
		fmt.Printf("Not enough arguments\n")
		fmt.Printf("Please use %s <base_price> <percentage> <years>\n", os.Args[0])
		fmt.Printf("For example increase 950 with 5%% for 9 years: %s 950 5 9\n", os.Args[0])
		os.Exit(1)
	}

	inputNumber = os.Args[1]
	inputPercentage = os.Args[2]
	inputPeriod = os.Args[3]

	result, err = strconv.ParseFloat(inputNumber, 64)
	if err != nil {
		fmt.Printf("error converting number string to float64\n%s\n", err.Error())
		os.Exit(1)
	}

	initialPercentage, err = strconv.ParseFloat(inputPercentage, 64)
	if err != nil {
		fmt.Printf("error converting percentage string to float64\n%s\n", err.Error())
		os.Exit(1)
	}

	period, err = strconv.ParseInt(inputPeriod, 10, 64)
	if err != nil {
		fmt.Printf("error converting period string to int\n%s\n", err.Error())
		os.Exit(1)
	}

	percentage = (100 + initialPercentage) / 100
	fmt.Printf("A %.0f percentage increase per year for %d years would result in:\n", initialPercentage, period)
	fmt.Printf("\tFirst year %.0f\n", result)
	for i = 1; i < period; i++ {
		result = result * percentage
		fmt.Printf("\tYear %d %.2f\n", i+1, result)
	}

}
