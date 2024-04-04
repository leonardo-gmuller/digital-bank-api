package postgres

type AccountsRepository struct {
	*Client
}

func NewAccountsRepository(client *Client) *AccountsRepository {
	return &AccountsRepository{client}
}
