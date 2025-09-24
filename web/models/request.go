package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint64 `json:"id"`
	Nickname    string `json:"nickname"`
	AuthorityId uint64 `json:"authorityId"`
	jwt.StandardClaims
}
