package main

import (
	"log"

	"github.com/kaito2/nand2tetris/internal"
)

func main() {
	inputFilename := "sample-data/FunctionCalls/SimpleFunction/SimpleFunction.vm"
	parser, err := internal.NewParser(inputFilename)
	if err != nil {
		log.Fatalf("Failed to get NewParser: %v", err)
	}

	outputFilename := "sample-data/FunctionCalls/SimpleFunction/SimpleFunction.asm"
	err = parser.Parse(outputFilename)
	if err != nil {
		log.Fatalf("Failed to Parse: %v", err)
	}
}
