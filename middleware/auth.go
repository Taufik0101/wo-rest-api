package middleware

import (
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Auth(jwtService service.JWTService, userServ service.UserService, redisServ service.RedisService) gin.HandlerFunc {
	return func(context *gin.Context) {
		newToken := service.ExtrakToken(context)
		token, errTok := jwtService.ValidateToken(newToken)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["ID"]
		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			//ID := claims["ID"]
			checkRedis := redisServ.CheckValueKey(ID.(string), newToken)
			if checkRedis {
				SplitID := strings.Split(ID.(string), "_")
				ids, _ := strconv.ParseUint(SplitID[1], 10, 32)
				us := userServ.FindByID(uint32(ids))
				context.Set("UserLog", us)
			}else {
				resp := helper.BuildErrorResponse("Token Invalid", "Key Or Value Not Found", nil)
				context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			}
		} else {
			redisServ.ClearToken(ID.(string))
			resp := helper.BuildErrorResponse("Invalid Token", errTok.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		}
	}
}

func CheckRole(jwtService service.JWTService, userServ service.UserService, redisServ service.RedisService, roles string) gin.HandlerFunc {
	return func(context *gin.Context) {
		newToken := service.ExtrakToken(context)
		token, errTok := jwtService.ValidateToken(newToken)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["ID"]
		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			//ID := claims["ID"]
			checkRedis := redisServ.CheckValueKey(ID.(string), newToken)
			if checkRedis {
				role := claims["user_role"].(string)
				if role == roles {
					SplitID := strings.Split(ID.(string), "_")
					ids, _ := strconv.ParseUint(SplitID[1], 10, 32)
					us := userServ.FindByID(uint32(ids))
					context.Set("UserLog", us)
				}else {
					resp := helper.BuildErrorResponse("Role Invalid", "Kesalahan Role", nil)
					context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
				}
			}else {
				//redisServ.ClearToken(ID.(string))
				resp := helper.BuildErrorResponse("Token Invalid", "Key Or Value Not Found", nil)
				context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			}
		} else {
			redisServ.ClearToken(ID.(string))
			resp := helper.BuildErrorResponse("Token Invalid", errTok.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		}
	}
}