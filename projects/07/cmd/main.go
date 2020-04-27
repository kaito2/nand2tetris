package main

import (
	"log"

	"github.com/kaito2/nand2tetris/internal"
)

func main() {
	inputPath := "sample-data/FunctionCalls/FibonacciElement"
	aggregator, err := internal.NewAggregator(inputPath)
	if err != nil {
		log.Fatalf("Failed to get NewParser: %v", err)
	}

	outputFilePath := "sample-data/FunctionCalls/FibonacciElement/out.asm"
	err = aggregator.ParseAll(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to ParseALL: %v", err)
	}
}
