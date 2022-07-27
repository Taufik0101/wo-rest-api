package service

import (
	"fmt"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type JWTService interface {
	GenerateTokenAuth(user entity.User) string
	ValidateToken(token string) (*jwt.Token, error)
	GenerateTokenForget(user entity.User) string
}

type jwtService struct {
	JWTKey []byte
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer: "Yanto0123",
		JWTKey: getJWTKey(),
	}
}

func getJWTKey() []byte {
	secretKey := []byte(os.Getenv("JWTKey"))
	if len(secretKey) == 0 {
		secretKey = []byte("Yanto0123")
	}
	return secretKey
}

func (data *jwtService) GenerateTokenAuth(user entity.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["ID"] = "auth_" + strconv.Itoa(int(user.Id))
	claims["user_email"] = user.Email
	claims["user_role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString(data.JWTKey)
	if err != nil {
		panic(err)
	}
	return t
}

func (data *jwtService) GenerateTokenForget(user entity.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["ID"] = "forget_" + strconv.Itoa(int(user.Id))
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString(data.JWTKey)
	if err != nil {
		panic(err)
	}
	return t
}

func (data *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token Invalid %v", t_.Header["alg"])
		}
		return data.JWTKey, nil
	})
}

func ExtrakToken(context *gin.Context) string {
	Header := context.GetHeader("Authorization")
	if !strings.Contains(Header, "Bearer") {
		resp := helper.BuildErrorResponse("Failed", "Invalid", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return ""
	}
	newToken := ""
	Atoken := strings.Split(Header, " ")
	if len(Atoken) == 2 {
		newToken = Atoken[1]
	}
	return newToken
}
