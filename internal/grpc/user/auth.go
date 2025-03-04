package user

import (
	"context"
	"github.com/MikeThesol/proto/proto/user/gen"
)

type serverAPI struct {
	gen.UnimplementedUserServer
}

func (s *serverAPI) GetUserByEmail(ctx context.Context, req *gen.GetUserByEmailRequest) (*gen.GetUserByEmailResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetUserByID(ctx context.Context, req *gen.GetUserByIDRequest) (*gen.GetUserByIDResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, req *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	panic("implement me")
}
