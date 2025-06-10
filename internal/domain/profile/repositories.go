package profile

import (
	"context"

	"github.com/gorkagg10/lovify-user-service/internal/domain/oauth"
)

type UserRepository interface {
	CreateUserProfile(context.Context, *UserProfile) (string, error)
	ConnectWithMusicProvider(ctx context.Context, userID string) error
	//StoreMusicProviderToken(ctx context.Context, userID string, accessToken *oauth.Token) error
	//StoreMusicProviderData(ctx context.Context, userID string, musicProviderData *MusicProviderData) error
}

type MusicProviderRepository interface {
	GetTopTracks(token *oauth.Token) ([]Track, error)
	GetTopArtists(token *oauth.Token) ([]Artist, error)
}
