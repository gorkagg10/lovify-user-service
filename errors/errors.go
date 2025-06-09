package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrInvalidCreateUserRequestMsg = "INVALID_CREATE_USER_REQUEST"
	ErrInvalidUsernameMsg          = "INVALID_USERNAME"
	ErrInvalidBirthdayMsg          = "INVALID_BIRTHDAY"
	ErrInvalidGenderMsg            = "INVALID_GENDER"
	ErrInvalidSexualOrientationMsg = "INVALID_SEXUAL_ORIENTATION"
)

var (
	StatusInvalidCreateUserRequest = status.New(codes.InvalidArgument, ErrInvalidCreateUserRequestMsg)
	StatusInvalidUsername          = status.New(codes.InvalidArgument, ErrInvalidUsernameMsg)
	StatusInvalidBirthday          = status.New(codes.InvalidArgument, ErrInvalidBirthdayMsg)
	StatusInvalidGender            = status.New(codes.InvalidArgument, ErrInvalidGenderMsg)
	StatusInvalidSexualOrientation = status.New(codes.InvalidArgument, ErrInvalidSexualOrientationMsg)

	ErrInvalidCreateUserRequest = StatusInvalidCreateUserRequest.Err()
	ErrInvalidUsername          = StatusInvalidUsername.Err()
	ErrInvalidBirthday          = StatusInvalidBirthday.Err()
	ErrInvalidGender            = StatusInvalidGender.Err()
	ErrInvalidSexualOrientation = StatusInvalidSexualOrientation.Err()
)
