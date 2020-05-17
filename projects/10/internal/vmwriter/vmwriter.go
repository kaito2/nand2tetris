package vmwriter

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

type VMWriter interface{}

type VMWriterImpl struct {
	file io.StringWriter
}

func NewVMWriterImpl(filename string) (VMWriter, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open file.")
	}
	return VMWriterImpl{
		file: file,
	}, nil
}

func (v *VMWriterImpl) writeLine(line string) error {
	if _, err := v.file.WriteString(line); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func (v *VMWriterImpl) WritePush(segment Segment, index int) error {
	line := fmt.Sprintf("push %s %d\n", segment, index)
	return v.writeLine(line)
}

func (v *VMWriterImpl) WritePop(segment Segment, index int) error {
	line := fmt.Sprintf("pop %s %d\n", segment, index)
	return v.writeLine(line)
}

func (v *VMWriterImpl) WriteArithmetic(command ArithmeticCommand) error {
	line := fmt.Sprintf("%s\n", command)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeLabel(label string) error {
	line := fmt.Sprintf("label %s\n", label)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeGoto(label string) error {
	line := fmt.Sprintf("goto %s\n", label)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeIf(label string) error {
	line := fmt.Sprintf("if-goto %s\n", label)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeCall(name string, nArgs int) error {
	line := fmt.Sprintf("call %s %d\n", name, nArgs)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeFunction(name string, nLocals int) error {
	line := fmt.Sprintf("function %s %d\n", name, nLocals)
	return v.writeLine(line)
}

func (v *VMWriterImpl) writeReturn() error {
	line := "label\n"
	return v.writeLine(line)
}

func (v *VMWriterImpl) close() {
	// TODO: implement
	// StringWriter を指定しているので Close() がない…
	panic("not implemented")
}
