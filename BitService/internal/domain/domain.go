package domain

type GraphStore interface{}
type ObjectStore interface{}

type Domain struct {
	graphDb  GraphStore
	objectDb ObjectStore
}

func NewDomain(graphDb GraphStore, objectDb ObjectStore) *Domain {
	return &Domain{
		graphDb:  graphDb,
		objectDb: objectDb,
	}
}

func (d *Domain) CreateBit() error {
	return nil
}
