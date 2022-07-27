package file

import (
	"fmt"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

func ImageProduk(c *gin.Context, a service.ProdukService, produk_id uint32) *multipart.FileHeader {
	fileProduk, errProduk := c.FormFile("produk")
	if errProduk != nil || fileProduk.Size == 0 {
		return nil
	}else {
		res := a.ProdukById(produk_id)
		if res.Image == "" {
			return MoveFileProduk(c, fileProduk)
		}else {
			errUnlink := UnlinkProduk(res.Image)
			if errUnlink != nil {
				return nil
			}else {
				return MoveFileProduk(c, fileProduk)
			}
		}

	}
}

func MoveFileProduk(c *gin.Context, fileProduk *multipart.FileHeader) *multipart.FileHeader {
	pathProduk := fmt.Sprintf("upload/images/produk/%s", fileProduk.Filename)
	erProf := c.SaveUploadedFile(fileProduk, pathProduk)
	if erProf != nil {
		return nil
	}else {
		return fileProduk
	}
}

func UnlinkProduk(image string) error {
	_, err := os.Stat("./upload/images/produk/" + image)
	if os.IsNotExist(err) {
		return err
	}
	err = os.Remove("./upload/images/produk/" + image)
	if err != nil {
		return err
	}
	return nil
}

func SaveImageProduk(c *gin.Context) *multipart.FileHeader {
	fileProduk, errKTP := c.FormFile("produk")
	if errKTP != nil || fileProduk.Size == 0 {
		return nil
	}else {
		return MoveFileProduk(c, fileProduk)
	}
}
