package data

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type VendorController interface {
	VendorByCategory(ctx *gin.Context)
	VendorByCity(ctx *gin.Context)
	VendorByKategoryAndCity(ctx *gin.Context)
	AllVendor(ctx *gin.Context)
}

type vendorController struct {
	userService service.UserService
	jwtService service.JWTService
}

func (v vendorController) VendorByCategory(ctx *gin.Context) {
	category, err := strconv.ParseUint(ctx.Param("category"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter reuqest_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		res := v.userService.VendorByCategory(uint32(category))
		resp := helper.BuildResponse(true, "Get Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (v vendorController) VendorByCity(ctx *gin.Context) {
	kota, err := strconv.ParseUint(ctx.Param("kota"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter reuqest_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		res := v.userService.VendorByCity(uint32(kota))
		resp := helper.BuildResponse(true, "Get Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (v vendorController) VendorByKategoryAndCity(ctx *gin.Context) {
	KC := ctx.Param("KC")
	KCSplit := strings.Split(KC, "-")
	category, errKat := strconv.ParseUint(KCSplit[0], 10, 64)
	kota, errKot := strconv.ParseUint(KCSplit[1], 10, 64)
	if errKat != nil || errKot != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter reuqest_id yang ditemukan", errKot.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		res := v.userService.VendorByKategoryAndCity(uint32(category),uint32(kota))
		resp := helper.BuildResponse(true, "Get Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (v vendorController) AllVendor(ctx *gin.Context) {
	var vendors []entity.User = v.userService.AllVendor()
	resp := helper.BuildResponse(true, "Get Data Berhasil", vendors)
	ctx.JSON(http.StatusOK, resp)
}

func NewVendorController(userServ service.UserService, jwtServ service.JWTService) VendorController {
	return &vendorController{
		userService: userServ,
		jwtService: jwtServ,
	}
}