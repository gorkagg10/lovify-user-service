package profile

import "time"

type UserProfile struct {
	username                 string
	birthday                 time.Time
	gender                   string
	sexualOrientation        string
	description              string
	connectedToMusicProvider bool
	musicProviderInfo        *MusicProviderData
}

func (u *UserProfile) Username() string {
	return u.username
}

func (u *UserProfile) Birthday() time.Time {
	return u.birthday
}

func (u *UserProfile) Gender() string {
	return u.gender
}

func (u *UserProfile) SexualOrientation() string {
	return u.sexualOrientation
}

func (u *UserProfile) Description() string {
	return u.description
}

func (u *UserProfile) MusicProviderInfo() *MusicProviderData {
	return u.musicProviderInfo
}

func (u *UserProfile) ConnectedToMusicProvider() bool {
	return u.connectedToMusicProvider
}

func NewUserProfile(
	username string,
	birthday time.Time,
	gender string,
	sexualOrientation string,
	description string,
) *UserProfile {
	return &UserProfile{
		username:                 username,
		birthday:                 birthday,
		gender:                   gender,
		sexualOrientation:        sexualOrientation,
		description:              description,
		connectedToMusicProvider: false,
	}
}
