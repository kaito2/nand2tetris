package internal

import (
	"fmt"
	"os"
)

type CodeWriter struct {
	// outputFilename string
	outputFile        *os.File
	currentVMFilename string
}

func NewCodeWriter(outputFilename string) (CodeWriter, error) {
	f, err := os.Create(outputFilename)
	if err != nil {
		return CodeWriter{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	return CodeWriter{
		outputFile: f,
		// currentVMFilename: "",
	}, nil
}

func (c *CodeWriter) setFilename(filename string) {
	c.currentVMFilename = filename
}

func (c CodeWriter) writeArithmetic(cmd string) {
	// convert to assemble code
	if isBinaryFunction(cmd) {
		c.writeBinaryFunction(cmd)
	} else { // isUnaryFunction
		c.writeUnaryFunction(cmd)
	}
}

func (c CodeWriter) writeBinaryFunction(cmd string) {
	switch symbol(cmd) {
	case eqSymbol, gtSymbol, ltSymbol:
		c.writeCompAndJumpFunction(cmd)
	default:
		c.writeDestAndCompFunction(cmd)
	}
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
	c.outputFile.WriteString("D=M-D\n")
	c.outputFile.WriteString("@TRUE\n")
	operator := getOperator(cmd)
	c.outputFile.WriteString(fmt.Sprintf("M;%s\n", operator))

	// FALSE
	c.outputFile.WriteString("@0\n")
	c.outputFile.WriteString("D=A\n")
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("M=D\n")

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M+1\n")
	c.outputFile.WriteString("@END\n")
	c.outputFile.WriteString("D;JMP\n")

	// TRUE
	c.outputFile.WriteString("(TRUE)\n")
	c.outputFile.WriteString("@65535\n")
	c.outputFile.WriteString("D=A\n")
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("M=D\n")

	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M+1\n")
	c.outputFile.WriteString("@END\n")
	c.outputFile.WriteString("D;JMP\n")

	c.outputFile.WriteString("(END)\n")

}

func (c CodeWriter) writeCompAndJumpFunction(cmd string) {

}

func (c CodeWriter) writeUnaryFunction(cmd string) {
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

func (c CodeWriter) writePushPop(cmdType CommandType, segment string, index uint16) {
	// TODO: validation
	if cmdType == C_PUSH {
		// XXX: support only constant segment
		// TODO: define segment enum
		if segment == "constant" {
			// TODO: avoid hard coding...
			// TODO: error handling
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("D=A\n")
			c.outputFile.WriteString("@SP\n")
			c.outputFile.WriteString("A=M\n")
			c.outputFile.WriteString("M=D\n")

			c.outputFile.WriteString("@SP\n")
			c.outputFile.WriteString("M=M+1\n")
		} else {
			// TODO: implement
			panic("not implemented !")
		}
	} else { // C_POP
		// TODO: implement
		panic("not implemented !")
	}
}
