package Service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(email string,role string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthCustomClaim struct {
	Name string `json:"name"`
	Role string `json:"role"`
	User bool  `json:"user"`
	jwt.StandardClaims
}

type JwtServices struct {
	SecretKey string
	Issuer	string
}

func JwtAuthService() *JwtServices {
	return &JwtServices{
		SecretKey: getSecretKey(),
		Issuer:    "FoodOrder API",
	}}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secretbangetdah"
	}
	return secret
}

func (service *JwtServices) GenerateToken(email string,role string, isUser bool) string  {
	claims := AuthCustomClaim{
		Name:           email,
		Role: 			role,
		User:           isUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*48).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    service.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.SecretKey))
	if err != nil{
		panic(err)
	}

	return t
}

func (service *JwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid{
			return nil, fmt.Errorf("Invalid Token ", token.Header["alg"])
		}

		return []byte(service.SecretKey), nil
	})
}

