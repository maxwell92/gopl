package sorter

type IResource interface {
	Print()
}

type LessFunc func(d1, d2 IResource) bool

type Sorter struct {
	elems []IResource
	Lessf []LessFunc
}
