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

type KategoriController interface {
	AllKategori(ctx *gin.Context)
	SimpanKategori(ctx *gin.Context)
	UpdateKategori(ctx *gin.Context)
	DeleteKategori(ctx *gin.Context)
}

type kategoriController struct {
	jwtService service.JWTService
	kategoriService service.CategoriService
}

func (k kategoriController) AllKategori(ctx *gin.Context) {
	var kategoris []entity.Categories = k.kategoriService.AllKategori()
	resp := helper.BuildResponse(true, "Get Data Berhasil", kategoris)
	ctx.JSON(http.StatusOK, resp)
}

func (k kategoriController) SimpanKategori(ctx *gin.Context) {
	var CKategori dto.CreateKategori
	errCreate := ctx.ShouldBind(&CKategori)
	if errCreate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		res := k.kategoriService.SimpanKategori(CKategori)
		resp := helper.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (k kategoriController) UpdateKategori(ctx *gin.Context) {
	kategori_id, err := strconv.ParseUint(ctx.Param("kategori_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter kategori_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var kategori dto.UpdateKategori
		errUpdate := ctx.ShouldBind(&kategori)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			res := k.kategoriService.UpdateKategori(uint32(kategori_id), kategori)
			resp := helper.BuildResponse(true, "Update Data Berhasil", res)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (k kategoriController) DeleteKategori(ctx *gin.Context) {
	kategori_id, err := strconv.ParseUint(ctx.Param("kategori_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter kategori_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		k.kategoriService.HapusKategori(uint32(kategori_id))
		resp := helper.BuildResponse(true, "Hapus Data Berhasil", nil)
		ctx.JSON(http.StatusOK, resp)
	}
}

func NewKategoriController(jwtServ service.JWTService, kategoriServ service.CategoriService) KategoriController {
	return &kategoriController{
		jwtService: jwtServ,
		kategoriService: kategoriServ,
	}
}
