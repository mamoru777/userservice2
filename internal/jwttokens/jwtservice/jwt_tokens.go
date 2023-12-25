package jwtservice

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mamoru777/userservice2/internal/mylogger"

	"time"
)

const (
	AccessTokenDuration  = time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	SecretKey            = "mamoru" // Замените на ваш секретный ключ
)

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime, expOk := claims["exp"].(float64)
		if !expOk {
			mylogger.Logger.Error("Неверный токен: пропущено 'exp' claim")
			err = errors.New("Неверный токен: пропущено 'exp' claim")
			return nil, err
		}

		// Преобразование времени из float64 в тип time.Time
		expiration := time.Unix(int64(expirationTime), 0)

		// Проверка срока годности токена
		if time.Now().After(expiration) {
			mylogger.Logger.Println("У токена вышел срок годности")
			err = errors.New("У токена вышел срок годности")
			return nil, err
		}
		return claims, nil
	}

	mylogger.Logger.Error("Неверный токен")
	err = errors.New("Неверный токен")
	return nil, err
}
