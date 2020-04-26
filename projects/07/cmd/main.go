package main

import (
	"log"

	"github.com/kaito2/nand2tetris/internal"
)

func main() {
	inputFilename := "sample-data/MemoryAccess/StaticTest/StaticTest.vm"
	parser, err := internal.NewParser(inputFilename)
	if err != nil {
		log.Fatalf("Failed to get NewParser: %v", err)
	}

	outputFilename := "sample-data/MemoryAccess/StaticTest/StaticTest.asm"
	err = parser.Parse(outputFilename)
	if err != nil {
		log.Fatalf("Failed to Parse: %v", err)
	}
}
