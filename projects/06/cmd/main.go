package main

import "github.com/kaito2/nand2tetris/internal"

func main() {
	filename := "sample-data/max/MaxL.asm"
	internal.Parse(filename)
}
