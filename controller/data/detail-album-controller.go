package data

import (
	"github.com/Taufik0101/wo-rest-api/controller/file"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DetailAlbumController interface {
	AllDetailAlbum(ctx *gin.Context)
	DetailAlbumByAlbum(ctx *gin.Context)
	SimpanDetailAlbum(ctx *gin.Context)
	HapusDetailAl(ctx *gin.Context)
}

type detailAlbumController struct {
	jwtService service.JWTService
	detailAlbumService service.DetailAlbumService
}

func (d detailAlbumController) AllDetailAlbum(ctx *gin.Context) {
	var detail_albums []entity.Detail_Album = d.detailAlbumService.AllDetailAlbum()
	resp := helper.BuildResponse(true, "Get Data Berhasil", detail_albums)
	ctx.JSON(http.StatusOK, resp)
}

func (d detailAlbumController) DetailAlbumByAlbum(ctx *gin.Context) {
	album_id, err := strconv.ParseUint(ctx.Param("album_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter album_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		us := d.detailAlbumService.DetailAlbumByAlbum(uint32(album_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", us)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (d detailAlbumController) SimpanDetailAlbum(ctx *gin.Context) {
	var CDAlbum dto.CreateDetailAlbum
	errUpdate := ctx.ShouldBind(&CDAlbum)
	if errUpdate != nil {
		response := helper.BuildErrorResponse("Failed to parsing", errUpdate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		save := file.SaveImageAlbum(ctx)
		CDAlbum.Foto = save.Filename
		res := d.detailAlbumService.SimpanDetailAlbum(CDAlbum)
		resp := helper.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (d detailAlbumController) HapusDetailAl(ctx *gin.Context) {
	detail_id, err := strconv.ParseUint(ctx.Param("detail_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter detail_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		find := d.detailAlbumService.DetailById(uint32(detail_id))
		errDel := file.UnlinkAlbum(find.Foto)
		if errDel != nil {
			resp := helper.BuildErrorResponse("Hapus Gagal", "Kesalahan menghapus gambar", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		}
		d.detailAlbumService.HapusDetail(uint32(detail_id))
		resp := helper.BuildResponse(true, "Hapus Data Berhasil", nil)
		ctx.JSON(http.StatusOK, resp)
	}
}

func NewDetailAlbumController(jwtServ service.JWTService, detalAlbumServ service.DetailAlbumService) DetailAlbumController {
	return &detailAlbumController{
		jwtService: jwtServ,
		detailAlbumService: detalAlbumServ,
	}
}
