package vmwriter

type VMWriter interface{}

type VMWriterImpl struct{}

func NewVMWriterImpl() (VMWriter, error) {
	return VMWriterImpl{}, nil
}

func (v *VMWriterImpl) writePush(segment Segment, index int) {
	panic("not implemented")
}

func (v *VMWriterImpl) writePop(segment Segment, index int) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeArithmetic(command ArithmeticCommand) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeLabel(label string) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeGoto(label string) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeIf(label string) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeCall(name string, nArgs int) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeFunction(name string, nLocals int) {
	panic("not implemented")
}

func (v *VMWriterImpl) writeReturn() {
	panic("not implemented")
}

func (v *VMWriterImpl) close() {
	panic("not implemented")
}
