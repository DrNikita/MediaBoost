package model

type Bit struct {
	Id       string
	AuthorId int
	Name     string
	Length   int
	Path     string
	Tags     []string
}

type BitTree struct {
	Bit
	Childs []BitTree
}
