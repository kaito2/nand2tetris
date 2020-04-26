package internal

import (
	"fmt"
	"os"
)

type CodeWriter struct {
	// outputFilename string
	outputFile        *os.File
	currentVMFilename string
	writeNum          uint16
}

func NewCodeWriter(outputFilename string) (CodeWriter, error) {
	f, err := os.Create(outputFilename)
	if err != nil {
		return CodeWriter{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	return CodeWriter{
		outputFile: f,
		// currentVMFilename: "",
		writeNum: 0,
	}, nil
}

func (c *CodeWriter) setFilename(filename string) {
	c.currentVMFilename = filename
}

func (c *CodeWriter) incWriteNum() {
	c.writeNum++
}

func (c *CodeWriter) writeArithmetic(cmd string) {
	defer c.incWriteNum()
	// convert to assemble code
	if isBinaryFunction(cmd) {
		c.writeBinaryFunction(cmd)
	} else { // isUnaryFunction
		c.writeUnaryFunction(cmd)
	}
}

func (c CodeWriter) writeBinaryFunction(cmd string) {
	c.outputFile.WriteString(fmt.Sprintf("// writeBinaryFunction (cmd: %s)\n", cmd))

	switch symbol(cmd) {
	case eqSymbol, gtSymbol, ltSymbol:
		c.writeCompAndJumpFunction(cmd)
	default:
		c.writeDestAndCompFunction(cmd)
	}
}

func (c CodeWriter) writeCompAndJumpFunction(cmd string) {
	// TODO: error handling
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("D=M\n")

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("D=M-D\n")
	trueSymbol := fmt.Sprintf("TRUE%d", c.writeNum)
	c.outputFile.WriteString(fmt.Sprintf("@%s\n", trueSymbol))
	operator := getOperator(cmd)
	c.outputFile.WriteString(fmt.Sprintf("D;%s\n", operator))

	// FALSE
	c.outputFile.WriteString("@0\n")
	c.outputFile.WriteString("D=A\n")
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("M=D\n")

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("D=M+1\n")
	endSymbol := fmt.Sprintf("END%d", c.writeNum)
	c.outputFile.WriteString(fmt.Sprintf("@%s\n", endSymbol))
	c.outputFile.WriteString("D;JMP\n")

	// TRUE
	c.outputFile.WriteString(fmt.Sprintf("(%s)\n", trueSymbol))
	c.outputFile.WriteString("@1\n")
	c.outputFile.WriteString("D=-A\n")
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("M=D\n")
	c.outputFile.WriteString(fmt.Sprintf("(%s)\n", endSymbol))

	// ポインタを進める
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M+1\n")
}

func (c CodeWriter) writeDestAndCompFunction(cmd string) {
	// TODO: error handling
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("D=M\n")

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")

	operator := getOperator(cmd)
	c.outputFile.WriteString(fmt.Sprintf("M=M%sD\n", operator))

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M+1\n")
}

func (c CodeWriter) writeUnaryFunction(cmd string) {
	c.outputFile.WriteString(fmt.Sprintf("// writeUnaryFunction(cmd: %s)\n", cmd))
	// TODO: error handling
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")

	// TODO: implement 'eq', 'gt', 'lt'
	operator := getOperator(cmd)
	c.outputFile.WriteString(fmt.Sprintf("M=%sM\n", operator))

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M+1\n")
}

func (c *CodeWriter) writePushPop(cmdType CommandType, segment string, index uint16) {
	c.outputFile.WriteString(fmt.Sprintf("// writePushPop(cmd: %v, segment: %s, index: %d)\n", cmdType, segment, index))
	defer c.incWriteNum()
	// TODO: validation
	if cmdType == C_PUSH {
		// XXX: support only constant segment
		// TODO: define segment enum
		if segment == "constant" {
			// TODO: error handling
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("D=A\n")

		} else if segment == "static" {
			c.outputFile.WriteString(fmt.Sprintf("@%s.%d\n", c.currentVMFilename, index))
			c.outputFile.WriteString("D=M\n")
		} else if segment == "temp" || segment == "pointer" {
			// TODO: implement
			c.outputFile.WriteString(fmt.Sprintf("@%s\n", getSegmentSymbol(segment)))
			// temp なら RAM[5 + index], pointer なら RAM[3 + index] になる必要があるので、
			// D=M ではなく D=A とする。
			c.outputFile.WriteString("D=A\n")
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("A=D+A\n")
			c.outputFile.WriteString("D=M\n")
		} else {
			// TODO: implement
			c.outputFile.WriteString(fmt.Sprintf("@%s\n", getSegmentSymbol(segment)))
			c.outputFile.WriteString("D=M\n")
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("A=D+A\n")
			c.outputFile.WriteString("D=M\n")
		}
		c.outputFile.WriteString("@SP\n")
		c.outputFile.WriteString("A=M\n")
		c.outputFile.WriteString("M=D\n")

		c.outputFile.WriteString("@SP\n")
		c.outputFile.WriteString("M=M+1\n")
	} else { // C_POP
		// TODO: validation (e.g. if segment == 'constant' then cause error)

		if segment == "temp" || segment == "pointer" {
			c.outputFile.WriteString(fmt.Sprintf("@%s\n", getSegmentSymbol(segment)))
			c.outputFile.WriteString("D=A\n")
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("D=D+A\n")
		} else if segment == "static" {
			c.outputFile.WriteString(fmt.Sprintf("@%s.%d\n", c.currentVMFilename, index))
			c.outputFile.WriteString("D=A\n")
		} else {
			c.outputFile.WriteString(fmt.Sprintf("@%s\n", getSegmentSymbol(segment)))
			c.outputFile.WriteString("D=M\n")
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("D=D+A\n")
		}
		// dest{c.writeNum} に対象アドレスを格納
		// OPTIMIZE: RAM 使うのはズル??
		c.outputFile.WriteString(fmt.Sprintf("@dest%d\n", c.writeNum))
		c.outputFile.WriteString("M=D\n")

		// pop
		c.outputFile.WriteString("@SP\n")
		c.outputFile.WriteString("M=M-1\n")
		c.outputFile.WriteString("A=M\n")
		c.outputFile.WriteString("D=M\n")

		// RAM[dest{c.writeNum}] に値を格納
		c.outputFile.WriteString(fmt.Sprintf("@dest%d\n", c.writeNum))
		c.outputFile.WriteString("A=M\n")
		c.outputFile.WriteString("M=D\n")
	}
}
