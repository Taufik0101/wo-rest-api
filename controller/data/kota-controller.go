package data

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type KotaController interface {
	AllKota(ctx *gin.Context)
	SimpanKota(ctx *gin.Context)
	UpdateKota(ctx *gin.Context)
	DeleteKota(ctx *gin.Context)
}

type kotaController struct {
	jwtService service.JWTService
	kotaService service.KotaService
}

func (k kotaController) AllKota(ctx *gin.Context) {
	var kotas []entity.Kota = k.kotaService.AllKota()
	resp := helper.BuildResponse(true, "Get Data Berhasil", kotas)
	ctx.JSON(http.StatusOK, resp)
}

func (k kotaController) SimpanKota(ctx *gin.Context) {
	var CKota dto.CreateKota
	errCreate := ctx.ShouldBind(&CKota)
	if errCreate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		res := k.kotaService.SimpanKota(CKota)
		resp := helper.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (k kotaController) UpdateKota(ctx *gin.Context) {
	kota_id, err := strconv.ParseUint(ctx.Param("kota_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter kota_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var kota dto.UpdateKota
		errUpdate := ctx.ShouldBind(&kota)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			res := k.kotaService.UpdateKota(uint32(kota_id), kota)
			resp := helper.BuildResponse(true, "Update Data Berhasil", res)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (k kotaController) DeleteKota(ctx *gin.Context) {
	kota_id, err := strconv.ParseUint(ctx.Param("kota_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter kota_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		k.kotaService.DeleteKota(uint32(kota_id))
		resp := helper.BuildResponse(true, "Hapus Data Berhasil", nil)
		ctx.JSON(http.StatusOK, resp)
	}
}

func NewKotaController(jwtServ service.JWTService, kotaServ service.KotaService) KotaController {
	return &kotaController{
		jwtService: jwtServ,
		kotaService: kotaServ,
	}
}