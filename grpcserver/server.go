package server

import (
	"almostintegration/pkg/grpcusers"
	"almostintegration/pkg/users"
	"context"
)

type server struct {
	fetcher users.Application
	grpcusers.UnimplementedUserGetServer
}

func (s *server) Users(context.Context, *grpcusers.EmptyParams) (*grpcusers.UsersReply, error) {
	uu, err := s.fetcher.Users()
	if err != nil {
		return nil, err
	}
	res := &grpcusers.UsersReply{
		Users: localUsersToGrpc(uu),
	}
	return res, nil
}

func localUsersToGrpc(uu []users.User) []*grpcusers.User {
	res := make([]*grpcusers.User, len(uu))
	for i := range uu {
		res[i] = &grpcusers.User{
			Name: uu[i].Name,
			Age:  int32(uu[i].Age),
		}
	}
	return res
}
