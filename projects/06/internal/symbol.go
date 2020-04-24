package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var preDefinedSymbolTable = map[string]uint16{
	"SP":     uint16(0),
	"LCL":    uint16(1),
	"ARG":    uint16(2),
	"THIS":   uint16(3),
	"THAT":   uint16(4),
	"R0":     uint16(0),
	"R1":     uint16(1),
	"R2":     uint16(2),
	"R3":     uint16(3),
	"R4":     uint16(4),
	"R5":     uint16(5),
	"R6":     uint16(6),
	"R7":     uint16(7),
	"R8":     uint16(8),
	"R9":     uint16(9),
	"R10":    uint16(10),
	"R11":    uint16(11),
	"R12":    uint16(12),
	"R13":    uint16(13),
	"R14":    uint16(14),
	"R15":    uint16(15),
	"SCREEN": uint16(16384),
	"KBD":    uint16(24576),
}

type SymbolParser struct {
	filepath    string
	romCounter  uint16
	symbolTable map[string]uint16
}

func NewSymbolParser(filepath string) SymbolParser {
	return SymbolParser{
		filepath:    filepath,
		symbolTable: preDefinedSymbolTable,
	}
}

func (s *SymbolParser) ScanLCommands() error {
	s.romCounter = 0
	s.symbolTable = preDefinedSymbolTable

	f, err := os.Open(s.filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	// var romCounter uint16 = 0
	for buf.Scan() {
		cmd := removeWhitespace(buf.Text())
		if cmd == "" {
			continue
		}
		switch commandType(cmd) {
		case LCommand:
			symbol := symbol(cmd)
			if err := s.addSymbol(symbol, s.romCounter); err != nil {
				return fmt.Errorf("failed to s.addSymbol: %w", err)
			}
		default:
			if !isSkipped(cmd) {
				s.romCounter++
			}
		}
	}
	return nil
}

func (s *SymbolParser) addSymbol(symbol string, address uint16) error {
	_, ok := s.symbolTable[symbol]
	if ok {
		return errors.New(fmt.Sprintf("symbol '%s' is already defined.", symbol))
	}
	s.symbolTable[symbol] = s.romCounter
	return nil
}

func (s SymbolParser) getAddress(symbol string) (uint16, bool) {
	address, ok := s.symbolTable[symbol]
	return address, ok
}
