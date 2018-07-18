package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]ObjData)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]ObjData
	outer *Environment
}

type ObjData struct {
	Obj      Object
	IsAssign bool
}

func (e *Environment) Get(name string) (ObjData, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val ObjData) ObjData {
	e.store[name] = val
	return val
}
