package mongodb

type UserProfile struct {
	ID                     string `bson:"_id,omitempty"`
	Username               string `bson:"username"`
	Birthday               string `bson:"birthday"`
	Gender                 string `bson:"gender"`
	SexualOrientation      string `bson:"sexual_orientation"`
	Description            string `bson:"description"`
	MusicProviderConnected bool   `bson:"music_provider_connected"`
}

func NewUserProfile(
	id string,
	username string,
	birthday string,
	gender string,
	sexualOrientation string,
	description string,
	musicProviderConnected bool,
) *UserProfile {
	return &UserProfile{
		ID:                     id,
		Username:               username,
		Birthday:               birthday,
		Gender:                 gender,
		SexualOrientation:      sexualOrientation,
		Description:            description,
		MusicProviderConnected: musicProviderConnected,
	}
}
