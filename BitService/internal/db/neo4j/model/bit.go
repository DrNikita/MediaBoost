package model

type Bit struct {
	Id            string
	AuthorId      int
	CoAuthorIds   []int
	Name          string
	Length        int
	Path          string
	Tags          []Tag
	AditionalTags []string
}

type BitTree struct {
	Bit
	Childs []BitTree
}
