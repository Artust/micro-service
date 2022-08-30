package jwt

import (
	"avatar/services/account_management/domain/entity"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId   int64
	Email    string
	Gender   int64
	RoleId   int64
	CenterId int64
	jwt.StandardClaims
}

func CreateJWT(acc *entity.Account, secretKey string, jwtExpirationHour int) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = acc.Id
	claims["email"] = acc.Email
	claims["gender"] = acc.Gender
	claims["roleId"] = acc.RoleId
	claims["centerId"] = acc.CenterId
	claims["displayName"] = acc.Username
	claims["avatar"] = acc.Avatar
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtExpirationHour)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string, secretKey string) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	jwtKey := []byte(secretKey)

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("ErrSignatureInvalid: ", err)
			return nil, err
		}
		fmt.Println("Err: ", err)
		return nil, err
	}
	if !tkn.Valid {
		fmt.Println("tkn.InValid: ", err)
		return nil, err
	}
	return claims, nil
}

func RefreshToken(token string, secretKey string, jwtExpirationHour int) (string, error) {
	claims, err := VerifyToken(token, secretKey)
	if err != nil {
		return "", err
	}
	return CreateJWT(&entity.Account{
		Id:       claims.UserId,
		Email:    claims.Email,
		Gender:   claims.Gender,
		RoleId:   claims.RoleId,
		CenterId: claims.CenterId,
	}, secretKey, jwtExpirationHour)
}
