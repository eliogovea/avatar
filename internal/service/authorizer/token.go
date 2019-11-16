package authorizer

import (
  "time"
  "errors"

  "github.com/dgrijalva/jwt-go"
)

var ErrUnauthorized = errors.New("unauthorized")

// TODO set an enviroment variable
var jwtKey = []byte("blablabla")

type Claims struct {
  Username string `json:"username"`
  Role     string `json:"role"`
  jwt.StandardClaims
}

func GenerateToken(username string, role string) (string) {
	claims := &Claims{
		Username: username,
    Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, _ := token.SignedString(jwtKey)
  return tokenString
}

func CheckToken(tokenString string) (error) {
  claims := &Claims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
  if err != nil || !token.Valid {
    return ErrUnauthorized
  }
  return nil
}

func GetUsername(tokenString string) (string, error) {
  claims := &Claims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
  if err != nil || !token.Valid {
    return "", ErrUnauthorized
  }
  return claims.Username, nil
}
