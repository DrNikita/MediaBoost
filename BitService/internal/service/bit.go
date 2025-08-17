package service

type GraphStore interface{}
type ObjectStore interface{}

type BitService struct {
	graphDb  GraphStore
	objectDb ObjectStore
}

func NewBitService(graphDb GraphStore, objectDb ObjectStore) *BitService {
	return &BitService{
		graphDb:  graphDb,
		objectDb: objectDb,
	}
}

func (b *BitService) DoSmthng() error {
	return nil
}
