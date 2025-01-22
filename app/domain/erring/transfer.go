package erring

var (
	ErrTransferAccountDestinationNotFound = NewAppError("transfer:account-destination-not-found", "account destination not found.")
	ErrTransferUserNotFound               = NewAppError("transfer:user-not-found", "user not found.")
	ErrTransferBalanceNotSufficient       = NewAppError("transfer:balance-not-sufficient", "the originating account does not have sufficient balance.")
)
