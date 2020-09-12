package jwts

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rijine/ads-api/internal/config"
	"time"
)

type JwtService interface {
	Generate(username string) (string, int64, error)
	Validate(token string) (string, error)
}

type service struct {
	issuer string
	secret string
	expiry time.Duration
}

type Claims struct {
	Username string `json:"username"`
	// TODO: ADD more
	jwt.StandardClaims
}

func NewJwtService(cfg config.JwtConfig) JwtService {
	return &service{
		expiry: cfg.Expiry,
		issuer: cfg.Issuer,
		secret: cfg.Secret,
	}
}

func (s *service) Generate(username string) (string, int64, error) {
	expiry := time.Now().Add(time.Hour * s.expiry).Unix()
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", expiry, err
	}

	return signed, expiry, nil

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
