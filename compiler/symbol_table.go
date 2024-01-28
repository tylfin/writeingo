package compiler

type SymbolScope string

const (
	GlobalScope   SymbolScope = "GLOBAL"
	LocalScope    SymbolScope = "LOCAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FunctionScope"
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
	FreeSymbols    []Symbol
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s, scope: GlobalScope, FreeSymbols: []Symbol{}}
}

func (st *SymbolTable) Define(ident string) Symbol {
	st.store[ident] = Symbol{Name: ident, Scope: st.scope, Index: st.numDefinitions}
	st.numDefinitions += 1
	return st.store[ident]
}

func (st *SymbolTable) DefineFunctionName(ident string) Symbol {
	st.store[ident] = Symbol{Name: ident, Scope: FunctionScope, Index: 0}
	return st.store[ident]
}

func (st *SymbolTable) defineFree(original Symbol) Symbol {
	st.FreeSymbols = append(st.FreeSymbols, original)

	symbol := Symbol{Name: original.Name, Index: len(st.FreeSymbols) - 1}
	symbol.Scope = FreeScope
	st.store[original.Name] = symbol

	return symbol
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]
	if !ok && st.parent != nil {
		obj, ok = st.parent.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := st.defineFree(obj)
		return free, true
	}

	return obj, ok
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
