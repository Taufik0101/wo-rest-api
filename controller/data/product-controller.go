package data

import (
	"encoding/json"
	file2 "github.com/Taufik0101/wo-rest-api/controller/file"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProdukController interface {
	AllProduk(ctx *gin.Context)
	ProdukByVendor(ctx *gin.Context)
	SimpanProduk(ctx *gin.Context)
	UpdateProduk(ctx *gin.Context)
	DeleteProduk(ctx *gin.Context)
	RandomProductLimit(ctx *gin.Context)
	PublishProduct(ctx *gin.Context)
	ProdukByVendorUser(ctx *gin.Context)
}

type produkController struct {
	jwtService service.JWTService
	produkService service.ProdukService
}

func (p produkController) ProdukByVendorUser(ctx *gin.Context) {
	vendor_id, err := strconv.ParseUint(ctx.Param("vendor_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter user_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var produks []entity.Product = p.produkService.ProdukByVendor(uint32(vendor_id), "User")
		resp := helper.BuildResponse(true, "Get Data Berhasil", produks)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (p produkController) PublishProduct(ctx *gin.Context) {
	produk_id, err := strconv.ParseUint(ctx.Param("produk_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter produk_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		jsonData, errJSON := ioutil.ReadAll(ctx.Request.Body)
		if errJSON != nil {

		}
		var dataForProduk map[string]string
		json.Unmarshal([]byte(jsonData), &dataForProduk)
		produkNow := p.produkService.ProdukById(uint32(produk_id))
		var UProduk dto.UpdateProduk
		UProduk.Nama_Produk = produkNow.Nama_Produk
		UProduk.Image = produkNow.Image
		UProduk.Seller = produkNow.Seller
		UProduk.DP = produkNow.DP
		UProduk.Harga = produkNow.Harga
		UProduk.Deskripsi = produkNow.Deskripsi
		UProduk.Is_Active = dataForProduk["status"]

		us := p.produkService.UpdateProduk(uint32(produk_id), UProduk)
		resp := helper.BuildResponse(true, "Update Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (p produkController) AllProduk(ctx *gin.Context) {
	var produks []entity.Product = p.produkService.AllProduk()
	resp := helper.BuildResponse(true, "Get Data Berhasil", produks)
	ctx.JSON(http.StatusOK, resp)
}

func (p produkController) ProdukByVendor(ctx *gin.Context) {
	vendor_id, err := strconv.ParseUint(ctx.Param("vendor_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter user_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var produks []entity.Product = p.produkService.ProdukByVendor(uint32(vendor_id), "Vendor")
		resp := helper.BuildResponse(true, "Get Data Berhasil", produks)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (p produkController) SimpanProduk(ctx *gin.Context) {
	var CProduk dto.CreateProduk
	errCreate := ctx.ShouldBind(&CProduk)
	if errCreate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		saveProd := file2.SaveImageProduk(ctx)
		if saveProd == nil {
			resp := helper.BuildErrorResponse("Register Gagal", "Kesalahan menyimpan gambar", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		}else {
			CProduk.Image = saveProd.Filename
			res := p.produkService.SimpanProduk(CProduk)
			resp := helper.BuildResponse(true, "Register Berhasil", res)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (p produkController) UpdateProduk(ctx *gin.Context) {
	produk_id, err := strconv.ParseUint(ctx.Param("produk_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter produk_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var UProduk dto.UpdateProduk
		errUpdate := ctx.ShouldBind(&UProduk)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			file := file2.ImageProduk(ctx, p.produkService, uint32(produk_id))
			if file != nil {
				UProduk.Image = file.Filename
			}
			us := p.produkService.UpdateProduk(uint32(produk_id), UProduk)
			resp := helper.BuildResponse(true, "Update Data Berhasil", us)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (p produkController) DeleteProduk(ctx *gin.Context) {
	produk_id, err := strconv.ParseUint(ctx.Param("produk_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter produk_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		prod := p.produkService.ProdukById(uint32(produk_id))
		errS := file2.UnlinkProduk(prod.Image)
		if errS != nil {
			resp := helper.BuildErrorResponse("Hapus Gagal", "Kesalahan menghapus gambar", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		}
		p.produkService.DeleteProduk(uint32(produk_id))
		resp := helper.BuildResponse(true, "Hapus Data Berhasil", nil)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (p produkController) RandomProductLimit(ctx *gin.Context) {
	var produks []entity.Product = p.produkService.RandomProductLimit()
	resp := helper.BuildResponse(true, "Get Data Berhasil", produks)
	ctx.JSON(http.StatusOK, resp)
}

func NewProdukController(jwtServ service.JWTService, produkServ service.ProdukService) ProdukController {
	return &produkController{
		jwtService: jwtServ,
		produkService: produkServ,
	}
}
