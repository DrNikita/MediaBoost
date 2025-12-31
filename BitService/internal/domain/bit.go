package domain

type Bit struct {
	Id            string
	AuthorId      int
	CoAuthorIds   []int
	Name          string
	Length        int
	ObjectPath    string
	Tags          []Tag
	AditionalTags []string
	LayerLevel    int
}

type BitTree struct {
	Bit
	Childs []BitTree
}
