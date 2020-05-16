package vmwriter

type Segment string

const (
	Const   Segment = "constant"
	Arg     Segment = "argument"
	Local   Segment = "local"
	Static  Segment = "static"
	This    Segment = "this"
	That    Segment = "that"
	Pointer Segment = "pointer"
	Temp    Segment = "temp"
)

type ArithmeticCommand string

const (
	Add ArithmeticCommand = "add"
	Sub ArithmeticCommand = "sub"
	Neg ArithmeticCommand = "neg"
	Eq  ArithmeticCommand = "eq"
	Gt  ArithmeticCommand = "gt"
	Lt  ArithmeticCommand = "lt"
	And ArithmeticCommand = "and"
	Or  ArithmeticCommand = "or"
	Not ArithmeticCommand = "not"
)
