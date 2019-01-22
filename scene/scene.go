package scene

type ElemRef interface {
}

type Scene struct {
	Id         string
	Name       string
	Persistent bool
	Element    []ElemRef
}
