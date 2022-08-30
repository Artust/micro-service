package jwt

import (
	"avatar/services/account_management/domain/entity"
	"fmt"
	"testing"
	"time"
)

const (
	JWT_SECRET_KEY      = "secret"
	JWT_EXPIRATION_HOUR = 12
)

func TestLoad(t *testing.T) {
	fmt.Println("SECRET: ", JWT_SECRET_KEY)
}

var Norman = &entity.Account{
	Id:        1,
	Email:     "test@check",
	Gender:    1,
	RoleId:    1,
	CenterId:  1,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: time.Time{},
}

func TestCreate(t *testing.T) {
	token, err := CreateJWT(Norman, JWT_SECRET_KEY, JWT_EXPIRATION_HOUR)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Token: ", token)
	claims, err := VerifyToken(token, JWT_SECRET_KEY)
	if err != nil {
		fmt.Println("Err verify: ", err)
	}
	fmt.Println("««««« Token is valid! »»»»»")
	fmt.Println("Claim user: ", claims.Email)
	time.Sleep(5 * time.Second)
	newToken, err := RefreshToken(token, JWT_SECRET_KEY, JWT_EXPIRATION_HOUR)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("New token: ", newToken)
}
