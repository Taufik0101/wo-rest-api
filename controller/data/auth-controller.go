package data

import (
	"encoding/json"
	"fmt"
	"github.com/Taufik0101/wo-rest-api/controller/file"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Email(ctx *gin.Context, req entity.Request) bool
	EmailTolak(ctx *gin.Context, email string) bool
	Logout(ctx *gin.Context)
	SendTokenForgot(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

type authController struct {
	redisService service.RedisService
	authService service.AuthService
	jwtService service.JWTService
	userService service.UserService
}

func (a authController) ResetPassword(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {

	}

	var dataForUser map[string]string
	json.Unmarshal([]byte(jsonData), &dataForUser)

	us := a.userService.FindByEmail(dataForUser["email"])
	data := us.(entity.User)

	if data.ForgotToken == dataForUser["otp"] {
		if time.Now().Before(data.ExpiretAt) {
			hashPW, _ := repository.Hash(dataForUser["password"])
			usss := a.userService.UpdateUser(data.Id, dto.Update{
				Password: string(hashPW),
				ForgotToken: "",
				ExpiretAt: time.Time{},
			})
			resp := helper.BuildResponse(true, "Ubah Password Berhasil", usss)
			ctx.JSON(http.StatusOK, resp)
		} else {
			response := helper.BuildErrorResponse("Reset Password Gagal", "Token Expired", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	} else {
		response := helper.BuildErrorResponse("Reset Password Gagal", "Token Tidak Sama", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (a authController) SendTokenForgot(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {

	}

	var dataForUser map[string]string
	json.Unmarshal([]byte(jsonData), &dataForUser)
	pass := file.RandomToken(4, 1, 1, 1);
	us := a.userService.FindByEmail(dataForUser["email"])
	data := us.(entity.User)
	timein := time.Now().Local().Add(time.Hour * time.Duration(0) +
		time.Minute * time.Duration(1) +
		time.Second * time.Duration(0))
	var dtUpdate dto.Update
	dtUpdate.ForgotToken = pass
	dtUpdate.ExpiretAt = timein

	var dtForgot dto.Forgot
	dtForgot.URL = "http://localhost:8080/forgotPassword/"+pass

	file.SendEmailForgot(dataForUser["email"], dtForgot)

	up := a.userService.UpdateUser(data.Id, dtUpdate)
	resp := helper.BuildResponse(true, "Update Data Berhasil", up)
	ctx.JSON(http.StatusOK, resp)
}

func (a authController) EmailTolak(ctx *gin.Context, email string) bool {
	file.SendEmailTolak(email)
	return true
}

func (a authController) Logout(ctx *gin.Context) {
	newToken := service.ExtrakToken(ctx)
	token, _ := a.jwtService.ValidateToken(newToken)
	claims := token.Claims.(jwt.MapClaims)
	a.redisService.ClearToken(claims["ID"].(string))
	response := helper.BuildResponse(true, "Logout Berhasil!", nil)
	ctx.JSON(http.StatusOK, response)
}

func (a authController) Email(ctx *gin.Context, req entity.Request) bool {
	pass := file.RandomPassword(8, 2, 2, 2);
	use := a.authService.Register(dto.Register{
		Name:     req.Name,
		Role:     req.Role,
		NoHP:     req.NoHP,
		NIK:      req.NIK,
		Email:    req.Email,
		Password: pass,
		Category: req.Category,
	})
	fmt.Println(pass)
	data := dto.Email{
		Email:    use.Email,
		Password: pass,
	}
	file.SendEmailVerification(use.Email, data)
	return true
}

func (a authController) Login(ctx *gin.Context) {
	var Login dto.Login
	errLog := ctx.ShouldBind(&Login)
	if errLog != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errLog.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authRes := a.authService.VerifyPassword(Login.Email, Login.Password)
	if v, ok := authRes.(entity.User); ok {
		generateToken := a.jwtService.GenerateTokenAuth(v)
		a.redisService.SaveToken("auth_" + strconv.Itoa(int(v.Id)), generateToken, 24 * time.Hour)
		v.Token = generateToken
		response := helper.BuildResponse(true, "Login Berhasil!", v)
		ctx.JSON(http.StatusOK, response)
	}else {
		response := helper.BuildErrorResponse("Login Gagal", "Email atau Password Salah", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (a authController) Register(ctx *gin.Context) {
	var Reg dto.Register
	errReg := ctx.ShouldBind(&Reg)
	if errReg != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errReg.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !a.authService.IsDuplicateEmail(Reg.Email) {
		response := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errReg.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
	} else {
		save := a.authService.Register(Reg)
		r := helper.BuildResponse(true, "OK", save)
		ctx.JSON(http.StatusCreated, r)
	}
}

func NewAuthController(authServ service.AuthService, jwtServ service.JWTService, redisServ service.RedisService, userServ service.UserService) AuthController {
	return &authController{
		authService: authServ,
		jwtService: jwtServ,
		redisService: redisServ,
		userService: userServ,
	}
}
