package authdto

type SignUpInputDto struct {
	UsernameAndPasswordDto
}

type SignUpOutputDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
