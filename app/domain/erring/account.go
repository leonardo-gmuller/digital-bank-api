package erring

var (
	ErrAccountExists       = NewAppError("account:exists", "account already exists.")
	ErrAccountCPFIsInvalid = NewAppError("account:cpf-is-invalid", "cpf is invalid.")
	ErrAccountNotExists    = NewAppError("account:not-exists", "account not exists.")
)
