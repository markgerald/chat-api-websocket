package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/models"
	"github.com/markgerald/chat-api-challenge/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Email string `json:"email"`
	Id    uint
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	var dbUser models.User
	dbUserNew := repository.GetUserByEmail(user.Email, dbUser)
	err := bcrypt.CompareHashAndPassword([]byte(dbUserNew.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{"message": "Incorrect password"})
		return
	}
	expirationTime := time.Now().Add(90 * time.Minute)
	claims := &Claims{
		Id:    dbUserNew.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	c.JSON(200, gin.H{"token": tokenString})
}
