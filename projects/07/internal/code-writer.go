package internal

import (
	"fmt"
	"log"
	"os"
)

type CodeWriter struct {
	// outputFilename string
	outputFile        *os.File
	currentVMFilename string
	writeNum          uint16
}

func NewCodeWriter(outputFilename string) (CodeWriter, error) {
	// NOTE: outputFile should be exists
	f, err := os.OpenFile(outputFilename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
	_, err := c.outputFile.WriteString("@SP\n")
	if err != nil {
		log.Fatalf("Fatal: %v\n", err)
	}
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
		} else if segment == "argument" || segment == "local" || segment == "this" || segment == "that" {
			// TODO: implement
			// index 分だけ A=A+1 し続ける作戦にする?
			c.outputFile.WriteString(fmt.Sprintf("@%s\n", getSegmentSymbol(segment)))
			c.outputFile.WriteString("D=M\n")
			c.outputFile.WriteString(fmt.Sprintf("@%d\n", index))
			c.outputFile.WriteString("A=D+A\n")
			c.outputFile.WriteString("D=M\n")
		} else {
			if index != 0 {
				panic("Symbol and index cannot be used at the same time.")
			}
			c.outputFile.WriteString(fmt.Sprintf("@%s.%s\n", c.currentVMFilename, segment))
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

func (c *CodeWriter) writeInit() {
	// TODO: implement
	panic("not implemented !")
}

func (c *CodeWriter) writeLabel(label string) {
	c.outputFile.WriteString(fmt.Sprintf("(%s.%s)\n", c.currentVMFilename, label))
}

func (c *CodeWriter) writeGoto(label string) {
	c.outputFile.WriteString(fmt.Sprintf("@%s.%s\n", c.currentVMFilename, label))
	c.outputFile.WriteString("0;JMP\n")
}

func (c *CodeWriter) writeIf(label string) {
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("D=M\n")

	c.outputFile.WriteString(fmt.Sprintf("@%s.%s\n", c.currentVMFilename, label))
	c.outputFile.WriteString("D;JNE\n")
}

func (c *CodeWriter) writeCall(functionName string, numArgs uint16) {
	// TODO: segment の代わりに label を使っても機能するか確認
	returnAddressSymbol := fmt.Sprintf("%s.return-address%d", c.currentVMFilename, c.writeNum)
	// push return-address
	c.writePushPop(C_PUSH, returnAddressSymbol, 0)
	// push LCL
	c.writePushPop(C_PUSH, "local", 0)
	// push ARG
	c.writePushPop(C_PUSH, "argument", 0)
	// push THIS
	c.writePushPop(C_PUSH, "this", 0)
	// push THAT
	c.writePushPop(C_PUSH, "that", 0)

	// ARG = SP - n - 5
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("D=M\n")
	for i := uint16(0); i < (numArgs + 5); i++ {
		c.outputFile.WriteString("D=D-1\n")
	}
	c.outputFile.WriteString("@ARG\n")
	c.outputFile.WriteString("M=D\n")

	// LCL = SP
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("D=M\n")
	c.outputFile.WriteString("@LCL\n")
	c.outputFile.WriteString("M=D\n")

	// goto f
	c.outputFile.WriteString(fmt.Sprintf("@%s.%s\n", c.currentVMFilename, functionName))
	c.outputFile.WriteString("0;JMP\n")
	c.outputFile.WriteString(fmt.Sprintf("(%s)\n", returnAddressSymbol))
}

func (c *CodeWriter) writeReturn() {
	// FRAME = LCL
	c.outputFile.WriteString("@LCL\n")
	c.outputFile.WriteString("D=M\n")
	c.outputFile.WriteString("@FRAME\n")
	c.outputFile.WriteString("M=D\n")

	// RET = *(FRAME - 5)
	c.writeLoad("RET", "FRAME", 5)

	// *ARG = pop()
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=M-1\n")
	c.outputFile.WriteString("A=M\n")
	c.outputFile.WriteString("D=M\n")

	c.outputFile.WriteString("@ARG\n")
	c.outputFile.WriteString("M=D\n")

	// SP = ARG + 1
	c.outputFile.WriteString("@ARG\n")
	c.outputFile.WriteString("D=M+1\n")
	c.outputFile.WriteString("@SP\n")
	c.outputFile.WriteString("M=D\n")

	// THAT = *(FRAME - 1)
	c.writeLoad("THAT", "FRAME", 1)
	// THIS = *(FRAME - 2)
	c.writeLoad("THIS", "FRAME", 2)
	// ARG  = *(FRAME - 3)
	c.writeLoad("ARG", "FRAME", 3)
	// LCL  = *(FRAME - 4)
	c.writeLoad("LCL", "FRAME", 4)

	// goto RET
	c.outputFile.WriteString("@RET\n")
	c.outputFile.WriteString("0;JMP\n")
}

// write assembly (pesudo: '[SYMBOL] = *(SEGMENT - n)')
// e.g. 'RET = *(FRAME - 5)'
func (c *CodeWriter) writeLoad(symbol, segment string, index uint16) {
	c.outputFile.WriteString(fmt.Sprintf("@%s\n", segment))
	c.outputFile.WriteString("D=M\n")

	//XXX: index を プラスしたい場合には対応できていない
	for i := uint16(0); i < index; i++ {
		c.outputFile.WriteString("D=D-1\n")
	}
	c.outputFile.WriteString("A=D\n")
	c.outputFile.WriteString("D=M\n")
	c.outputFile.WriteString(fmt.Sprintf("@%s\n", symbol))
	c.outputFile.WriteString("M=D\n")
}

func (c *CodeWriter) writeFunction(functionName string, numLocals uint16) {
	c.outputFile.WriteString(fmt.Sprintf("(%s.%s)\n", c.currentVMFilename, functionName))
	// ローカル変数の領域を確保
	for i := uint16(0); i < numLocals; i++ {
		c.writePushPop(C_PUSH, "constant", 0)
	}
}
