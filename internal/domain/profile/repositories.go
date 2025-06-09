package profile

import (
	"context"
)

type UserRepository interface {
	CreateUserProfile(context.Context, *UserProfile) (string, error)
}
