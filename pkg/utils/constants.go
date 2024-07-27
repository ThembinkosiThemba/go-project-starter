package utils

import "errors"

// ///////////////////////////////////////////////////
// //////////////////   ERRORS   /////////////////////
// ///////////////////////////////////////////////////
var (
	ErrExists                  = errors.New("user already exists")
	ErrUserNotFound            = errors.New("user not found")
	ErrFailedToCreate          = errors.New("failed to create user")
	ErrFailedToGetAllRecords   = errors.New("failed to get all records")
	ErrFailedToGetSingleRecord = errors.New("failed to get single record")
	ErrInternalServerError     = errors.New("something went wrong, internal server error")
	ErrFailedToDelete          = errors.New("failed to delete account")
)
