package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbolParser SymbolParser
var ramCount uint16

func Parse(filename, binFilename string) error {
	// ファイル用意
	binFile, err := os.Create(binFilename)
	if err != nil {
		return fmt.Errorf("failed to os.Open: %w", err)
	}
	defer binFile.Close()

	// ラベル (e.g. '(Xxx)') をスキャン
	symbolParser = NewSymbolParser(filename)
	if err := symbolParser.ScanLCommands(); err != nil {
		return fmt.Errorf("failed to ScanLCommand: %w", err)
	}

	// 命令をパース
	// RAMの15までは定義済みシンボル
	ramCount = 16
	parser, err := NewParser(filename)
	if err != nil {
		return fmt.Errorf("failed to get new parser: %w", err)
	}
	defer parser.close()
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
	commentRemovedCommand := removeComment(whitespaceRemovedCommand)
	switch commandType(commentRemovedCommand) {
	case ACommand:
		return parseACommand(commentRemovedCommand), true
	case CCommand:
		return parseCCommand(commentRemovedCommand), true
	default: // LCommand
		// do nothing when LCommand
		return 0, false
	}
}

func parseACommand(cmd string) uint16 {
	symbol := symbol(cmd)
	// TODO: 明示的にb15に0をセットしたほうが良い
	num, err := strconv.ParseUint(symbol, 10, 15)
	if err == nil {
		return uint16(num)
	}
	// TODO: symbolParser がグローバル変数になっているので修正
	address, ok := symbolParser.getAddress(symbol)
	if ok {
		return address
	} else {
		// TODO: error handling
		symbolParser.addSymbol(symbol, ramCount)
		ret := ramCount
		ramCount++
		return ret
	}
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

func removeComment(cmd string) string {
	return strings.Split(cmd, "//")[0]
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
