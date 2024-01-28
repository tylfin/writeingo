package compiler

type SymbolScope string

const (
	GlobalScope  SymbolScope = "GLOBAL"
	LocalScope   SymbolScope = "LOCAL"
	BuiltinScope SymbolScope = "BUILTIN"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
	scope          SymbolScope
	parent         *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s, scope: GlobalScope}
}

func (st *SymbolTable) Define(ident string) Symbol {
	st.store[ident] = Symbol{Name: ident, Scope: st.scope, Index: st.numDefinitions}
	st.numDefinitions += 1
	return st.store[ident]
}

func (st *SymbolTable) Resolve(ident string) (Symbol, bool) {
	if v, ok := st.store[ident]; ok || st.parent == nil {
		if !ok {
			return v, ok
		}
		return v, ok
	}

	return st.parent.Resolve(ident)
}

func (st *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Scope: BuiltinScope, Index: index}
	st.store[name] = symbol
	return symbol
}

func NewEnclosedSymbolTable(st *SymbolTable) *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s, scope: LocalScope, parent: st}
}
