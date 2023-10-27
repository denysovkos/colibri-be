package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = "SecretYouShouldHide" // os.Getenv("API_SECRET")

func GenerateHashedPassword(password string) string {
	bytesPassword := []byte("MyDarkSecret")

	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword[:])
}

func ComparePasswords(hashedPassword, password string) error {
	error := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if error != nil {
		return error
	}

	return nil
}

func GenerateToken(userId uint, userEmail string) (string, error) {
	token_lifespan, err := strconv.Atoi("1") // os.Getenv("TOKEN_HOUR_LIFESPAN")
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["userEmail"] = userEmail
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
