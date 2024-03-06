package authdto

import "github.com/golang-jwt/jwt/v5"

type JwtPayloadStruct struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
