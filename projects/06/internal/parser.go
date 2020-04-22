package internal

type CommandType int

const (
	ACommand CommandType = iota
	CCommand
	LCommand
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func hasMoreCommand() bool {
	panic("not implemented")
}

func advance() {
	panic("not implemented")
}

func commandType() CommandType {
	panic("not implemented")
}

func symbol() string {
	panic("not implemented")
}

func dest() string {
	panic("not implemented")
}

func comp() string {
	panic("not implemented")
}

func jump() string {
	panic("not implemented")
}
