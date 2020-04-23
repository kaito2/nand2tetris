package internal

import "strings"

type Mnemonic struct {
	String string
	Binary byte
}

var DestMap = map[string]uint16{
	"null": 0b000,
	"M":    0b001,
	"D":    0b010,
	"MD":   0b011,
	"A":    0b100,
	"AM":   0b101,
	"AD":   0b110,
	"AMD":  0b111,
}

func destBits(mnemonic string) uint16 {
	b, _ := DestMap[mnemonic]
	return b
}

var JumpMap = map[string]uint16{
	"null": 0b000,
	"JGT":  0b001,
	"JEQ":  0b010,
	"JGE":  0b011,
	"JLT":  0b100,
	"JNE":  0b101,
	"JLE":  0b110,
	"JMP":  0b111,
}

func jumpBits(mnemonic string) uint16 {
	b, _ := JumpMap[mnemonic]
	return b
}

// NOTE: before lookup, replace 'M' => 'A'
var CompMap = map[string]uint16{
	"0":   0b101010,
	"1":   0b111111,
	"-1":  0b111010,
	"D":   0b001100,
	"A":   0b110000,
	"!D":  0b001101,
	"!A":  0b110001,
	"-D":  0b001111,
	"-A":  0b110011,
	"D+1": 0b011111,
	"A+1": 0b110111,
	"D-1": 0b001110,
	"A-1": 0b110010,
	"D+A": 0b000010,
	"D-A": 0b010011,
	"A-D": 0b000111,
	"D&A": 0b000000,
	"D|A": 0b010101,
}

func compBits(mnemonic string) uint16 {
	m := mnemonic
	aBit := uint16(0b0)
	if strings.Contains(mnemonic, "M") {
		m = strings.Replace(mnemonic, "M", "A", 1)
		aBit = uint16(0b1 << 6)
	}
	b, _ := CompMap[m]
	return b | aBit
}
