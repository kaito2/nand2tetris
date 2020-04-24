package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename, binFilename string) error {
	parser, err := NewParser(filename)
	if err != nil {
		return fmt.Errorf("failed to get new parser: %w", err)
	}
	defer parser.close()

	binFile, err := os.Create(binFilename)
	if err != nil {
		return fmt.Errorf("failed to os.Open: %w", err)
	}
	defer binFile.Close()

	// 行を読み出してはパース
	for parser.advance() {
		b, incrementPC := parseCommand(parser.currentCommand)
		if !incrementPC {
			continue
		}
		_, err := binFile.WriteString(fmt.Sprintf("%016b\n", b))
		if err != nil {
			return err
		}
	}
	return nil
}

func parseCommand(cmd string) (uint16, bool) {
	whitespaceRemovedCommand := removeWhitespace(cmd)
	if isSkipped(whitespaceRemovedCommand) {
		return 0, false
	}
	switch commandType(whitespaceRemovedCommand) {
	case ACommand:
		return parseACommand(whitespaceRemovedCommand), true
	case CCommand:
		return parseCCommand(whitespaceRemovedCommand), true
	default: // LCommand
		// TODO: implement
		panic("not implemented")
	}
}

func parseACommand(cmd string) uint16 {
	symbol := symbol(cmd)
	// TODO: 明示的にb15に0をセットしたほうが良い
	num, err := strconv.ParseUint(symbol, 10, 15)
	if err == nil {
		return uint16(num)
	}
	// symbol is not digit
	// TODO: symbol を 数値に変更したものを返すように変更
	return uint16(num)
}

func parseCCommand(cmd string) uint16 {
	prefix := uint16(0b111 << 13)
	comp := compBits(comp(cmd)) << 6
	dest := destBits(dest(cmd)) << 3
	jump := jumpBits(jump(cmd))
	return prefix | comp | dest | jump
}

func removeWhitespace(line string) string {
	return strings.Replace(line, " ", "", -1)
}

func isSkipped(line string) bool {
	if len(line) == 0 {
		return true
	}
	// start with "//"
	if line[:2] == "//" {
		return true
	}
	return false
}
