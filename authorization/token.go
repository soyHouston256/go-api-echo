package authorization

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/soyhouston256/go-api/model"
	"time"
)

func GenerateToken(data *model.Login) (string, error) {
	claims := model.Claim{
		Email: data.Email,
		Name:  "SoyHouston",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "SoyHouston",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (model.Claim, error) {
	claims := &model.Claim{}
	token, err := jwt.ParseWithClaims(t, claims, verifyFunction)

	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("invalid token")
	}
	return *claims, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
