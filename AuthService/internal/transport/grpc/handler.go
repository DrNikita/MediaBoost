package grpc

import (
	"auth-service/internal/core/model"
	pb "auth-service/internal/transport/grpc/gen/auth_service"
	"context"
	"errors"
)

type Auth interface {
	SignUp(model.User)
	SingIn()
}

type Authenticator struct {
	pb.UnimplementedAuthServer
	auth Auth
}

func NewAuthenticator(auth Auth) Authenticator {
	return Authenticator{
		auth: auth,
	}
}

func (a Authenticator) SignUp(_ context.Context, userData *pb.UserData) (*pb.SignUpReply, error) {
	user := model.NewUser(int(userData.Age), userData.Role, userData.Email, userData.Nickname, userData.Password, userData.FirstName, userData.SecondName)
	a.auth.SignUp(user)
	return nil, errors.New("fffffffffffffffffffffffffffffffffffffffffffff")
}

func (a Authenticator) SignIn(context.Context, *pb.SignInCredentials) (*pb.SignInReply, error) {
	return nil, errors.New("fffffffffffffffffffffffffffffffffffffffffffff")
}
