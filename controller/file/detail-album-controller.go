package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

func MoveFileAlbum(c *gin.Context, fileAlbum *multipart.FileHeader) *multipart.FileHeader {
	pathAlbum := fmt.Sprintf("upload/images/album/%s", fileAlbum.Filename)
	erProf := c.SaveUploadedFile(fileAlbum, pathAlbum)
	if erProf != nil {
		return nil
	}else {
		return fileAlbum
	}
}

func UnlinkAlbum(image string) error {
	_, err := os.Stat("./upload/images/album/" + image)
	if os.IsNotExist(err) {
		return err
	}
	err = os.Remove("./upload/images/album/" + image)
	if err != nil {
		return err
	}
	return nil
}

func SaveImageAlbum(c *gin.Context) *multipart.FileHeader {
	fileAlbum, errKTP := c.FormFile("album")
	if errKTP != nil || fileAlbum.Size == 0 {
		return nil
	}else {
		return MoveFileAlbum(c, fileAlbum)
	}
}
