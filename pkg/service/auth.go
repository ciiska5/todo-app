package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/ciiska5/todo-app/entities"
	"github.com/ciiska5/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt           = "dgnajp]SFLKSADGJKSAGM,ldlfkmdsfm935mfskau3rqnfj" //дополнительная соль для хэширования пароля
	tokensValidDur = time.Hour * 12                                    //продолжительность действия токена JWT
	signingKey     = "kdshflksdf8940gjrsd3r039w890kgtfldsmjh2"         //ключ для шифровки/дешифровки токена JWT
)

// расширяем стандартные клеймы
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) CreateUser(user entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) GenerateToken(nickname, password string) (string, error) {
	user, err := as.repo.GetUser(nickname, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	//генерация токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokensValidDur).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// хэширование пароля пользователя
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
