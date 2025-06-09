package profile

import (
	"context"

	userServiceGrpc "github.com/gorkagg10/lovify-user-service/grpc/user-service"
)

type Manager struct {
	userRepository UserRepository
}

func NewManager(userRepository UserRepository) *Manager {
	return &Manager{userRepository: userRepository}
}

func (m *Manager) CreateUserProfile(ctx context.Context, req *userServiceGrpc.CreateUserRequest) (string, error) {
	userProfile := NewUserProfile(
		req.GetUsername(),
		req.GetBirthday().AsTime(),
		req.GetGender().String(),
		req.GetSexualOrientation().String(),
		req.GetDescription(),
	)
	return m.userRepository.CreateUserProfile(ctx, userProfile)
}
