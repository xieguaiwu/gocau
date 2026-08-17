package main

import (
	"fmt"
	"os"
)

const (
	MAX_NUMS   = 100
	maxFMetric = 120
)

var (
	command     string
	mutiSum     float64
	mutiNums    [MAX_NUMS]float64
	mutiWeights [MAX_NUMS]float64
	count       int
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  gocau avg    -> calculate average of numbers")
	fmt.Println("  gocau wavg   -> calculate weighted average of numbers")
	fmt.Println("  gocau fitness_metric --> calculate languistic Fitness Metric")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		Usage()
		return
	}
	command = args[1]
	handler()
}

func handler() {
	switch command {
	case "avg":
		readNumbers()
		fmt.Printf("Average = %.4f\n", calAvg())
	case "wavg":
		readWeightedNumbers()
		fmt.Printf("Weighted Average = %.4f\n", calWAvg())
	case "fitness_metric":
		readFitnessMetric()
		fmt.Printf("Fitness Metric = %.4f\n", calFMetric())
	default:
		Usage()
	}
}

func readNumbers() {
	fmt.Print("Enter number of values: ")
	fmt.Scan(&count)
	if count <= 0 || count > MAX_NUMS {
		fmt.Println("Invalid count")
		os.Exit(1)
	}
	mutiSum = 0
	fmt.Println("Enter all the numbers:")
	for i := 0; i < count; i++ {
		fmt.Scan(&mutiNums[i])
		mutiSum += mutiNums[i]
	}
}

func calAvg() float64 {
	return mutiSum / float64(count)
}

func readWeightedNumbers() {
	fmt.Print("Enter number of values: ")
	fmt.Scan(&count)
	if count <= 0 || count > MAX_NUMS {
		fmt.Println("Invalid count")
		os.Exit(1)
	}
	fmt.Println("Enter numbers: ")
	for i := 0; i < count; i++ {
		fmt.Scan(&mutiNums[i])
	}
	fmt.Println("Enter corresponding weights:")
	for i := 0; i < count; i++ {
		fmt.Scan(&mutiWeights[i])
	}
}

func calWAvg() float64 {
	var sumWeighted, sumWeights float64
	for i := 0; i < count; i++ { //weighted average number
		sumWeighted += mutiNums[i] * mutiWeights[i]
		sumWeights += mutiWeights[i]
	}
	if sumWeights == 0 {
		return 0
	}
	return sumWeighted / sumWeights
}

var FMetricTotal int
var FMetricScale float64
var FMetricConstantB float64
var FMetricError [maxFMetric]float64
var FMetricSupersetgrammar [maxFMetric]float64
var FMetricKnots [maxFMetric]float64
var FMetricSum [3]float64
var FMetericParserIndex int

func readFitnessMetric() {
	fmt.Printf("input total rounds of calculation > ")
	fmt.Scan(&FMetricTotal)
	FMetricScale = 100.0
	for FMetricScale >= 1 {
		fmt.Printf("\ninput the scaling factor (which is below one) > ")
		fmt.Scan(&FMetricScale)
		fmt.Printf("\n")
	}
	FMetricConstantB = -10
	for FMetricConstantB <= 1 {
		fmt.Printf("\ninput the superset penalty (which is larger than one) > ")
		fmt.Scan(&FMetricConstantB)
		fmt.Printf("\n")
	}
	for i := 0; i < FMetricTotal; i++ {
		fmt.Printf("Input the amount of violations at round %d: ", i+1)
		fmt.Scan(&FMetricError[i])
		fmt.Printf("Input the number of parameters set to superset values at round %d: ", i+1)
		fmt.Scan(&FMetricSupersetgrammar[i])
		fmt.Printf("Input the measure of the general elegance at round %d: ", i+1)
		fmt.Scan(&FMetricKnots[i])
	}
}

func calFMetric() float64 {
	FMetericParserIndex = -10
	for FMetericParserIndex < 0 || FMetericParserIndex > FMetricTotal {
		fmt.Printf("\ninput the parser index (which is above zero) > ")
		fmt.Scan(&FMetericParserIndex)
		fmt.Printf("\n")
	}
	//initialization
	FMetricSum[0] = 0
	FMetricSum[1] = 0
	FMetricSum[2] = 0

	for i := 0; i < FMetricTotal; i++ {
		FMetricSum[0] += FMetricError[i]
		FMetricSum[1] += FMetricSupersetgrammar[i]
		FMetricSum[2] += FMetricKnots[i]
	}
	mutiSum = FMetricSum[0] + FMetricConstantB*FMetricSum[1] + FMetricScale*FMetricSum[2]
	return (mutiSum - FMetricError[FMetericParserIndex] - FMetricSupersetgrammar[FMetericParserIndex] - FMetricKnots[FMetericParserIndex]) / mutiSum
}
