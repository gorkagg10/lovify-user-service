package profile

import (
	"context"
	"encoding/base64"
	userServiceGrpc "github.com/gorkagg10/lovify-user-service/grpc/user-service"
	"github.com/gorkagg10/lovify-user-service/internal/domain/oauth"
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

func (m *Manager) ConnectWithMusicProvider(ctx context.Context, state string, token *oauth.Token) error {
	userID, err := getUserID(state)
	if err != nil {
		return err
	}
	err = m.userRepository.ConnectWithMusicProvider(ctx, userID)
	if err != nil {
		return err
	}
	/*
		err = m.userRepository.StoreMusicProviderToken(ctx, userID, token)
		if err != nil {
			return err
		}
	*/
	return nil
}

func getUserID(state string) (string, error) {
	stateBytes, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return "", err
	}
	return string(stateBytes), nil
}
