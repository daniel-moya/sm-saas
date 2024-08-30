package auth

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your_secret_key")

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GenerateJWT(userID int64, userEmail string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "email": userEmail,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(jwtKey)
}
