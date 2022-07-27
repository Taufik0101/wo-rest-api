package data

import (
	file2 "github.com/Taufik0101/wo-rest-api/controller/file"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RequestController interface {
	AllRequest(ctx *gin.Context)
	SimpanRequest(ctx *gin.Context)
	UpdateRequest(ctx *gin.Context)
}

type requestController struct {
	jwtService service.JWTService
	requestService service.RequestService
	authControl AuthController
}

func (r requestController) AllRequest(ctx *gin.Context) {
	var requests []entity.Request = r.requestService.AllRequest()
	resp := helper.BuildResponse(true, "Get Data Berhasil", requests)
	ctx.JSON(http.StatusOK, resp)
}

func (r requestController) SimpanRequest(ctx *gin.Context) {
	var CRequest dto.CreateRequest
	errCreate := ctx.ShouldBind(&CRequest)
	if errCreate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		ktp, selfie := file2.SaveImageRequest(ctx)
		if ktp == "" || selfie == "" {
			resp := helper.BuildErrorResponse("Register Gagal", "Kesalahan menyimpan gambar", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		}else {
			CRequest.KTP = ktp
			CRequest.Selfie = selfie
			res := r.requestService.SimpanRequest(CRequest)
			resp := helper.BuildResponse(true, "Register Berhasil", res)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (r requestController) UpdateRequest(ctx *gin.Context) {
	req_id, err := strconv.ParseUint(ctx.Param("req_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter req_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var request dto.UpdateRequest
		errUpdate := ctx.ShouldBind(&request)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			if request.Status == "1" {
				res := r.requestService.UpdateRequest(uint32(req_id), request)
				if r.authControl.Email(ctx, res) {
					resp := helper.BuildResponse(true, "Data Berhasil Diterima", res)
					ctx.JSON(http.StatusOK, resp)
				}else {
					resp := helper.BuildErrorResponse("Kirim Email Gagal", "Ada Kesalahan", helper.EmptyObj{})
					ctx.JSON(http.StatusOK, resp)
				}
			}else {
				res := r.requestService.UpdateRequest(uint32(req_id), request)
				if r.authControl.EmailTolak(ctx, res.Email) {
					resp := helper.BuildResponse(true, "Data Berhasil Ditolak", res)
					ctx.JSON(http.StatusOK, resp)
				}else {
					resp := helper.BuildErrorResponse("Kirim Email Gagal", "Ada Kesalahan", helper.EmptyObj{})
					ctx.JSON(http.StatusOK, resp)
				}
			}
		}
	}
}

func NewRequestController(jwtServ service.JWTService, requestServ service.RequestService, authControl AuthController) RequestController {
	return &requestController{
		jwtService: jwtServ,
		requestService: requestServ,
		authControl: authControl,
	}
}
