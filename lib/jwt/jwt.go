package jwtlib

import (
	"github.com/golang-jwt/jwt/v5"
	authdto "kanlam/entity/dto/auth.dto"
	"kanlam/lib/config"
)

type StructJwtLib struct {
	signingKey    []byte
	signingMethod jwt.SigningMethod
}

var JwtLib = &StructJwtLib{
	signingMethod: jwt.SigningMethodHS256,
	signingKey:    []byte(config.GetJwtConfig().Secret),
}

func (r *StructJwtLib) createUserToken(payload authdto.JwtPayloadStruct) (string, error) {
	token := jwt.NewWithClaims(r.signingMethod, payload)
	ss, err := token.SignedString(r.signingKey)

	return ss, err
}
