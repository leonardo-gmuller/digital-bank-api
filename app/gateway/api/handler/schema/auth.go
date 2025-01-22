package schema

type AuthRequest struct {
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
