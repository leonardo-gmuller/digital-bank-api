package erring

var (
	ErrLoginUnauthorized    = NewAppError("login:unauthorized", "login unauthorized.")
	ErrLoginUserNotFound    = NewAppError("login:user-not-found", "user not found.")
	ErrLoginTokenNotCreated = NewAppError("login:token-not-created", "failed to create token.")
)
