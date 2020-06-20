package common

import "context"

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	//ContextUserID var
	ContextUserID = contextKey("userid")
	//ContextRoleID var
	ContextRoleID = contextKey("roleid")
	// ContextEmail var
	ContextEmail = contextKey("email")
	// ContextSignature var
	ContextSignature = contextKey("signature")
)

//GetUserIDFromContext helper get context value
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userid, ok := ctx.Value(ContextUserID).(string)
	return userid, ok
}

//GetSignatureFromContext helper get context value
func GetSignatureFromContext(ctx context.Context) (string, bool) {
	sign, ok := ctx.Value(ContextSignature).(string)
	return sign, ok
}

//GetEmailFromContext helper get email from context value
func GetEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(ContextEmail).(string)
	return email, ok
}

//GetRoleIDFromContext helper get roleid from context value
func GetRoleIDFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(ContextRoleID).(string)
	return role, ok
}
