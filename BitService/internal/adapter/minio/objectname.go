package minio

import "fmt"

const (
	bit = "bits"
)

type objectName struct {
	userId int
	rootId string
	name   string
}

func NewObjectName(name string, userId int, rootId ...string) *objectName {
	var objName string
	var objectName objectName
	if len(rootId) != 0 {
		objectName.rootId = rootId[0]
		objName = fmt.Sprintf("/%s/%d/%s/%s/", bit, userId, rootId[0], name)
	} else {
		objName = fmt.Sprintf("/%s/%d/%s/", bit, userId, name)
	}

	objectName.userId = userId
	objectName.name = objName
	return &objectName
}

func (o *objectName) Name() string {
	return o.name
}

func (o *objectName) GetBitPrefixForAuthor(authorId int) string {
	return fmt.Sprintf("/%s/%d/", bit, authorId)
}

func (o *objectName) GetBitPrefixForAuthorsBit(authorId int, bitId string) string {
	return fmt.Sprintf("/%s/%d/%s/", bit, authorId, bitId)
}
