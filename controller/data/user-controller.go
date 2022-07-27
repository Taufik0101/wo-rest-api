package data

import (
	file2 "github.com/Taufik0101/wo-rest-api/controller/file"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	AllUser(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
	kotaService service.KotaService
}

func (u userController) UpdatePassword(ctx *gin.Context) {
	user_id, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter user_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var UPassword dto.UPassword
		errUpdate := ctx.ShouldBind(&UPassword)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			Cek := u.userService.CekPassword(uint32(user_id), UPassword.OPassword)
			if !Cek {
				response := helper.BuildErrorResponse("Ubah Password gagal", "", helper.EmptyObj{})
				ctx.JSON(http.StatusBadRequest, response)
			}else {
				hashPW, _ := repository.Hash(UPassword.NPassword)
				us := u.userService.UpdateUser(uint32(user_id), dto.Update{Password: string(hashPW)})
				resp := helper.BuildResponse(true, "Ubah Password Berhasil", us)
				ctx.JSON(http.StatusOK, resp)
			}
		}
	}
}

func (u userController) AllUser(ctx *gin.Context) {
	var users []entity.User = u.userService.ListAll()
	resp := helper.BuildResponse(true, "Get Data Berhasil", users)
	ctx.JSON(http.StatusOK, resp)
}

func (u userController) FindByID(ctx *gin.Context) {
	user_id, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter user_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		us := u.userService.FindByID(uint32(user_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (u userController) UpdateUser(ctx *gin.Context) {
	user_id, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter user_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var userUpdateDTO dto.Update
		errUpdate := ctx.ShouldBind(&userUpdateDTO)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			file := file2.ImageProfile(ctx, u.userService, uint32(user_id))
			if file != nil {
				userUpdateDTO.Profil = file.Filename
			}
			if len(userUpdateDTO.KotaBaru) != 0 {
				if !u.kotaService.FindKotaByName(userUpdateDTO.KotaBaru) {
					response := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errUpdate.Error(), helper.EmptyObj{})
					ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
				} else {
					newKota := u.kotaService.SimpanKota(dto.CreateKota{Nama_Kota: userUpdateDTO.KotaBaru})
					userUpdateDTO.Kota = newKota.Id_Kota
				}
			}
			us := u.userService.UpdateUser(uint32(user_id), userUpdateDTO)
			resp := helper.BuildResponse(true, "Update Data Berhasil", us)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func NewUserController(userServ service.UserService, jwtServ service.JWTService, kotaServ service.KotaService) UserController {
	return &userController{
		userService: userServ,
		jwtService: jwtServ,
		kotaService: kotaServ,
	}
}
