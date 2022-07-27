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

type AlbumController interface {
	AllAlbum(ctx *gin.Context)
	SimpanAlbum(ctx *gin.Context)
	AlbumByVendor(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type albumController struct {
	jwtService service.JWTService
	albumService service.AlbumService
}

func (a albumController) Update(ctx *gin.Context) {
	album_id, err := strconv.ParseUint(ctx.Param("album_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter album_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		var album dto.UpdateAlbum
		errUpdate := ctx.ShouldBind(&album)
		if errUpdate != nil {
			response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
		}else {
			res := a.albumService.UpdateAlbum(uint32(album_id), album)
			resp := helper.BuildResponse(true, "Update Data Berhasil", res)
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func (a albumController) AllAlbum(ctx *gin.Context) {
	var albums []entity.Album = a.albumService.AllAlbum()
	resp := helper.BuildResponse(true, "Get Data Berhasil", albums)
	ctx.JSON(http.StatusOK, resp)
}

func (a albumController) SimpanAlbum(ctx *gin.Context) {
	var CAlbum dto.CreateAlbum
	errCreate := ctx.ShouldBind(&CAlbum)
	if errCreate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		res := a.albumService.SimpanAlbum(CAlbum)
		resp := helper.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (a albumController) AlbumByVendor(ctx *gin.Context) {
	vendor_id, err := strconv.ParseUint(ctx.Param("vendor_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter vendor_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		us := a.albumService.AlbumByVendor(uint32(vendor_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func NewAlbumController(jwtServ service.JWTService, albumServ service.AlbumService) AlbumController {
	return &albumController{
		jwtService: jwtServ,
		albumService: albumServ,
	}
}