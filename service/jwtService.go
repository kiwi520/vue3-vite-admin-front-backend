package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//var jwtKey = []byte("secret_key")
type JwtService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken()
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey []byte
	issuer    string
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{}
	claims.UserID = userID
	claims.ExpiresAt = time.Now().AddDate(0, 0, 1).Unix()
	claims.Issuer = j.issuer
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(j.secretKey)
	if err != nil {
		log.Println(err)
	}

	t, err := token.SignedString(signKey)

	if err != nil {
		panic(err)
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		verifyBytes, err := ioutil.ReadFile(os.Getenv("AccessTokenPublicKeyPath"))
		if err != nil {
			log.Fatalln("unable to read public key", "error", err)
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Fatalln("unable to parse public key", "error", err)
			return nil, err
		}
		println("verifyKey")
		println(verifyKey)
		println("verifyKey")
		return verifyKey, nil
	})
}

func (j *jwtService) RefreshToken() {
	panic("implement me")
}

func NewJwtService() JwtService {
	return &jwtService{
		issuer:    "sdfdsfsfs",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() []byte {

	secretKey, err := ioutil.ReadFile(os.Getenv("AccessTokenPrivateKeyPath"))
	if err != nil {
		log.Fatalf("unable to read private key", "error", err)
	}

	return secretKey
}
