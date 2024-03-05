package authdto

type SignInInputDto struct {
	UsernameAndPasswordDto
}

type SignInOutputDto struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"RefreshToken"`
}
