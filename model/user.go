package model

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	// User is -> (tb_user) of the database.
	User struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Claim .
	Claim struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		jwt.StandardClaims
	}
)
