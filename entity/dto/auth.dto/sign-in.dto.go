package authdto

type SignInInputDto struct {
	UsernameAndPasswordDto
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type SignInOutputDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token
}
