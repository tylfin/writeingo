package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func (st *SymbolTable) Define(ident string) Symbol {
	st.store[ident] = Symbol{Name: ident, Scope: GlobalScope, Index: st.numDefinitions}
	st.numDefinitions += 1
	return st.store[ident]
}

func (st *SymbolTable) Resolve(ident string) (Symbol, bool) {
	v, ok := st.store[ident]
	return v, ok
}
