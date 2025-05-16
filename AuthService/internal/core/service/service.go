package service

import "auth-service/internal/core/model"

type Authenticator interface{}

type UserStore interface{}

type CoreService struct {
	auth  Authenticator
	store UserStore
}

func NewCoreService(auth Authenticator, store UserStore) CoreService {
	return CoreService{
		auth:  auth,
		store: store,
	}
}

func (cs *CoreService) SignUp(user *model.User) error {
	return nil
}

func (cs *CoreService) SignIn(creds *model.Credentials) error {
	return nil
}
