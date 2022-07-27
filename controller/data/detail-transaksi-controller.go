package data

import (
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DetailTransaksiController interface {
	DetailTransaksiSeller(ctx *gin.Context)
	UpdateDetailTransaksi(ctx *gin.Context)
}

type detailTransaksiController struct {
	jwtService service.JWTService
	detailTransaksiService service.DetailTransaksiService
}

func (d detailTransaksiController) DetailTransaksiSeller(ctx *gin.Context) {
	vendor_id, err := strconv.ParseUint(ctx.Param("vendor_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter vendor_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		us := d.detailTransaksiService.DetailTransaksiSeller(uint32(vendor_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (d detailTransaksiController) UpdateDetailTransaksi(ctx *gin.Context) {
	panic("implement me")
}

func NewDetailTransaksiController(jwtServ service.JWTService, detailTransasiServ service.DetailTransaksiService) DetailTransaksiController {
	return &detailTransaksiController{
		jwtService: jwtServ,
		detailTransaksiService: detailTransasiServ,
	}
}
