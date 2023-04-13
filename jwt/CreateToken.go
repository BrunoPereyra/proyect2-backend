package jwt

import (
	"backend/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateToken(user models.UserModel) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error")
	}
	TOKENPASSWORD := os.Getenv("TOKENPASSWORD")
	claims := jwt.MapClaims{
		"_id":      user.ID,
		"nameuser": user.NameUser,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(TOKENPASSWORD))
	return signedToken, err

}
