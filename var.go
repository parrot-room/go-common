package common

import "errors"

var (
	//DATETIMEFORMAT date time for times format
	DATETIMEFORMAT = "20060102150405"
	//ErrUnAuthorized error authentification
	ErrUnAuthorized = errors.New("user unathorized")
)
