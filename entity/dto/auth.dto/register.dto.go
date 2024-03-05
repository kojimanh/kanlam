package authdto

type RegisterInputDto struct {
	Username string `json:"username",validate:"required"`
	Password string `json:"password",validate:"required"`
}

type RegisterOutputDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
