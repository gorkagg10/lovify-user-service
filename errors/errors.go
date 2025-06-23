package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrInvalidCreateUserRequestMsg = "INVALID_CREATE_USER_REQUEST"
	ErrInvalidEmailMsg             = "INVALID_EMAIL"
	ErrInvalidBirthdayMsg          = "INVALID_BIRTHDAY"
	ErrInvalidGenderMsg            = "INVALID_GENDER"
	ErrInvalidSexualOrientationMsg = "INVALID_SEXUAL_ORIENTATION"
)

var (
	StatusInvalidCreateUserRequest = status.New(codes.InvalidArgument, ErrInvalidCreateUserRequestMsg)
	StatusInvalidEmail             = status.New(codes.InvalidArgument, ErrInvalidEmailMsg)
	StatusInvalidBirthday          = status.New(codes.InvalidArgument, ErrInvalidBirthdayMsg)
	StatusInvalidGender            = status.New(codes.InvalidArgument, ErrInvalidGenderMsg)
	StatusInvalidSexualOrientation = status.New(codes.InvalidArgument, ErrInvalidSexualOrientationMsg)

	ErrInvalidCreateUserRequest = StatusInvalidCreateUserRequest.Err()
	ErrInvalidEmail             = StatusInvalidEmail.Err()
	ErrInvalidBirthday          = StatusInvalidBirthday.Err()
	ErrInvalidGender            = StatusInvalidGender.Err()
	ErrInvalidSexualOrientation = StatusInvalidSexualOrientation.Err()
)
