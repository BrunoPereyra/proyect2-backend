package jwt

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func parseToken(tokenString string) (*jwt.Token, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv error parseToken")
	}
	TOKENPASSWORD := os.Getenv("TOKENPASSWORD")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TOKENPASSWORD), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func ExtractDataFromToken(tokenString string) (string, error) {
	// Primero, parsea el token

	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	// Luego, accede a los claims del token para obtener los datos que necesitas
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid claims")
	}
	nameUser, ok := claims["nameuser"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid nameUser")
	}
	return nameUser, nil
}
