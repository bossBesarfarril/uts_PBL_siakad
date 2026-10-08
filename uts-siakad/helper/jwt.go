package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, role string) (string, int, error) {
	expiresIn := 3600 // 1 Jam = 3600 detik

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, expiresIn, nil
}
