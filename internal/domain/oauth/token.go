package oauth

import "time"

type Token struct {
	accessToken  string
	refreshToken string
	expiresAt    time.Time
}

func NewToken(accessToken, refreshToken string, expiresAt time.Time) *Token {
	return &Token{
		accessToken:  accessToken,
		refreshToken: refreshToken,
		expiresAt:    expiresAt,
	}
}
