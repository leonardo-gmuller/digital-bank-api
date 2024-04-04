package postgres

type TransfersRepository struct {
	*Client
}

func NewTransfersRepository(client *Client) *TransfersRepository {
	return &TransfersRepository{client}
}
