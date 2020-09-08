package authorization

import (
	"errors"
	"fmt"
	"time"

	"github.com/Leonardo-Antonio/golang-echo/model"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken .
func GenerateToken(data *model.User) (string, error) {
	claim := model.Claim{
		ID:    data.ID,
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "CMCLeo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		fmt.Println("hols")
		return "", err
	}
	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("Token no valido")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("No se puedo obtener los claims")
	}
	return *claim, nil
}

func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
