package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorkagg10/lovify-user-service/internal/domain/profile"
)

type UserRepository struct {
	UserProfileCollection *mongo.Collection
}

func NewUserRepository(userProfileCollection *mongo.Collection) *UserRepository {
	return &UserRepository{
		UserProfileCollection: userProfileCollection,
	}
}

func (u *UserRepository) CreateUserProfile(ctx context.Context, profile *profile.UserProfile) (string, error) {
	userProfileID := uuid.New().String()
	userProfile := NewUserProfile(
		userProfileID,
		profile.Username(),
		profile.Birthday().String(),
		profile.Gender(),
		profile.SexualOrientation(),
		profile.Description(),
		profile.MusicProviderInfo().Connected(),
	)
	_, err := u.UserProfileCollection.InsertOne(ctx, userProfile)
	if err != nil {
		return "", err
	}
	return userProfileID, nil
}

func (u *UserRepository) ConnectWithMusicProvider(ctx context.Context, userID string) error {
	filter := bson.D{{"_id", userID}}
	update := bson.D{{"$set", bson.D{{"music_provider_connected", true}}}}
	_, err := u.UserProfileCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
