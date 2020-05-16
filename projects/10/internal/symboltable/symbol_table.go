package symboltable

import "github.com/kaito2/nand2tetris/internal/types"

type SymbolTable interface{}

type SymbolTableImpl struct {
	classScopeTable      Table
	subroutineScopeTable Table
}

func NewSymbolTable() (SymbolTable, error) {
	return SymbolTableImpl{}, nil
}

func (s *SymbolTableImpl) startSubroutine() {
	// initialize subroutine table
	s.subroutineScopeTable = Table{}
}

func (s *SymbolTableImpl) getScopedTable(kind types.Ident) Table {
	switch checkScope(kind) {
	case Subroutine:
		return s.subroutineScopeTable
	default: // Class
		return s.classScopeTable
	}
}

func (s *SymbolTableImpl) contains(name string) bool {
	if _, ok := s.subroutineScopeTable.table[name]; ok {
		return true
	}
	if _, ok := s.classScopeTable.table[name]; ok {
		return true
	}
	return false
}

func (s *SymbolTableImpl) define(name, identType string, kind types.Ident) {
	table := s.getScopedTable(kind)
	table.define(name, identType, kind)
	return
}

func (s *SymbolTableImpl) varCount(kind types.Ident) int {
	table := s.getScopedTable(kind)
	return table.varCount(kind)
}

func (s *SymbolTableImpl) typeOf(name string) string {
	if s.subroutineScopeTable.contains(name) {
		return s.subroutineScopeTable.typeOf(name)
	} else { // if s.classScopeTable.contains(name)
		return s.classScopeTable.typeOf(name)
	}
}

// nil を返すためにポインタにしている…
func (s *SymbolTableImpl) indexOf(name string) *int {
	var index int
	if s.subroutineScopeTable.contains(name) {
		index = s.subroutineScopeTable.indexOf(name)
	} else if s.classScopeTable.contains(name) {
		index = s.classScopeTable.indexOf(name)
	}
	return &index
}

type Scope int

const (
	Class Scope = iota
	Subroutine
)

func checkScope(kind types.Ident) Scope {
	switch kind {
	case types.Static, types.Field:
		return Class
	case types.Argument, types.Var:
		return Subroutine
	default:
		panic("unknown ident type.")
	}
}

type Symbol struct {
	index      int
	name       string
	symbolType string
	kind       types.Ident
}

type table map[string]Symbol
type Table struct {
	table
	nextIndex int
}

func (t *Table) define(name, symbolType string, kind types.Ident) {
	// FIXME: デフォルトで上書きする
	t.table[name] = Symbol{
		name:       name,
		symbolType: symbolType,
		kind:       kind,
		index:      t.nextIndex,
	}
	t.nextIndex++
	return
}

func (t *Table) varCount(kind types.Ident) int {
	cnt := 0
	for _, symbol := range t.table {
		if symbol.kind == kind {
			cnt++
		}
	}
	return cnt
}

func (t Table) contains(name string) bool {
	_, ok := t.table[name]
	return ok
}

func (t Table) kindOf(name string) types.Ident {
	return t.table[name].kind
}

func (t Table) typeOf(name string) string {
	return t.table[name].symbolType
}

func (t Table) indexOf(name string) int {
	return t.table[name].index
}
