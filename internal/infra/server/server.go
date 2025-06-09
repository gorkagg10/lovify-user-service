package server

import (
	userServiceGrpc "lovify-user-service/grpc/user-service"
)

type UserServer struct {
	userServiceGrpc.UnimplementedUserServiceServer
}
