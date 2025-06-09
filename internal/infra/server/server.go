package server

import (
	userServiceGrpc "github.com/gorkagg10/lovify-user-service/grpc/user-service"
)

type UserServer struct {
	userServiceGrpc.UnimplementedUserServiceServer
}
