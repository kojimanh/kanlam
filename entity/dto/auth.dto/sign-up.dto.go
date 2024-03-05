package authdto

type SignUpInputDto struct {
	Username string `json:"username",validate:"required"`
	Password string `json:"password",validate:"required"`
}

type SignUpOutputDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
