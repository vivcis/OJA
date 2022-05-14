package services

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const AccessTokenValidity = time.Minute * 20
const RefreshTokenValidity = time.Hour * 24

var decode = []byte(os.Getenv("DECODE_SECRET_KEY"))

type Claims struct {
	UserEmail string `json:"email"`
	jwt.StandardClaims
}

// GetTokenFromHeader returns the token string in the authorization header
func GetTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) > 8 {
		return authHeader[7:]
	}
	return ""
}

// verifyAccessToken verifies a token
func verifyToken(tokenString *string, claims jwt.MapClaims, secret *string) (*jwt.Token, error) {
	parser := &jwt.Parser{SkipClaimsValidation: true}
	return parser.ParseWithClaims(*tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*secret), nil
	})
}

// AuthorizeToken check if a refresh token is valid
func AuthorizeToken(token *string, secret *string) (*jwt.Token, jwt.MapClaims, error) {
	if token != nil && *token != "" && secret != nil && *secret != "" {
		claims := jwt.MapClaims{}
		token, err := verifyToken(token, claims, secret)
		if err != nil {
			return nil, nil, err
		}
		return token, claims, nil
	}
	return nil, nil, fmt.Errorf("empty token or secret")
}

// GenerateToken generates only an access token
func GenerateToken(signMethod *jwt.SigningMethodHMAC, claims jwt.MapClaims, secret *string) (*string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(signMethod, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(*secret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *Service) GenerateNonAuthToken(UserEmail string, secret string) (*string, error) {
	// Define expiration time
	expirationTime := time.Now().Add(60 * time.Minute)
	// define the payload with the expiration time
	claims := &Claims{
		UserEmail: UserEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token with secret key
	tokenString, err := token.SignedString([]byte(secret))
	log.Println(tokenString)
	return &tokenString, err

}
func (s *Service) DecodeToken(token, secret string) (string, error) {

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return claims.UserEmail, err
}
