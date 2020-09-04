package jwts

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtService interface {
	Generate(username string) (string, error)
	Validate(token string) (string, error)
}

type service struct {
	issuer string
	secret string
	expiry time.Duration
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJwtService(issuer, secret string, expiry time.Duration) JwtService {
	return &service{
		expiry: expiry,
		issuer: issuer,
		secret: secret,
	}
}

func (s *service) Generate(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * s.expiry).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return signed, nil

}

func (s *service) Validate(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexepected signing method %s", t.Header["alg"])
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(*Claims); ok {
		if t.Valid && claims.ExpiresAt > time.Now().Unix() {
			fmt.Println("is valid")
			return claims.Username, nil
		}
		/* TODO: Need to verify
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", ve
			}
		}*/
	}

	return "", err

	/*return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexepected signing method %s", t.Header["alg"])
		}
		return []byte(s.secret), nil
	})*/
}
