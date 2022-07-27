package data

import (
	"encoding/json"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PenilaianController interface {
	AllPenilaian(ctx *gin.Context)
	PenilaianByVendor(ctx *gin.Context)
	SimpanPenilaian(ctx *gin.Context)
}

type penilaianController struct {
	jwtService service.JWTService
	penilaianService service.PenilaianService
	detailTransaksiService service.DetailTransaksiService
	transaksiService service.TransaksiService
}

func (p penilaianController) AllPenilaian(ctx *gin.Context) {
	var penilaians []entity.Penilaian = p.penilaianService.AllPenilaian()
	resp := helper.BuildResponse(true, "Get Data Berhasil", penilaians)
	ctx.JSON(http.StatusOK, resp)
}

func (p penilaianController) PenilaianByVendor(ctx *gin.Context) {
	vendor_id, err := strconv.ParseUint(ctx.Param("vendor_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter vendor_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		us := p.penilaianService.PenilaianByVendor(uint32(vendor_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (p penilaianController) SimpanPenilaian(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {

	}
	var dataPenilaian map[string]string
	json.Unmarshal([]byte(jsonData), &dataPenilaian)

	id_transaksi, _ := strconv.ParseUint(dataPenilaian["id_transaksi"], 10, 64)
	star, _ := strconv.ParseUint(dataPenilaian["star"], 10, 64)

	detail := p.detailTransaksiService.DetailTransaksi(uint32(id_transaksi))

	for _, value := range detail{
		var CPenilaian dto.CreatePenilaian

		CPenilaian.Customer = value.Customer
		CPenilaian.Seller = value.Seller
		CPenilaian.Pesan = dataPenilaian["pesan"]
		CPenilaian.Judul = dataPenilaian["judul"]
		CPenilaian.Star = uint16(star)
		CPenilaian.Trans = value.IdTransaksi

		_ = p.penilaianService.SimpanPenilaian(CPenilaian)
	}

	var UTransaksi dto.UpdateTransaksi

	UTransaksi.IsReview = dataPenilaian["star"]

	_ = p.transaksiService.UpdateTransaksi(uint32(id_transaksi), UTransaksi)

	resp := helper.BuildResponse(true, "Review Telah Ditambahkan", nil)
	ctx.JSON(http.StatusOK, resp)
	//var CPenilaian dto.CreatePenilaian
	//errCreate := ctx.ShouldBind(&CPenilaian)
	//if errCreate != nil {
	//	response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
	//	ctx.JSON(http.StatusBadRequest, response)
	//}else {
	//	res := p.penilaianService.SimpanPenilaian(CPenilaian)
	//	resp := helper.BuildResponse(true, "Tambah Data Berhasil", res)
	//	ctx.JSON(http.StatusOK, resp)
	//}
}

func NewPenilaianController(jwtServ service.JWTService, penilaianServ service.PenilaianService, detailTransaksiServ service.DetailTransaksiService, transaksiServ service.TransaksiService) PenilaianController {
	return &penilaianController{
		jwtService: jwtServ,
		penilaianService: penilaianServ,
		detailTransaksiService: detailTransaksiServ,
		transaksiService: transaksiServ,
	}
}
