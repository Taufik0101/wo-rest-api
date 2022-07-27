package file

import (
	"fmt"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

func ImageProfile(c *gin.Context, a service.UserService, user_id uint32) *multipart.FileHeader {
	fileProfil, errKTP := c.FormFile("profile")
	if errKTP != nil || fileProfil.Size == 0 {
		return nil
	}else {
		res := a.FindByID(user_id)
		if res.Profil == "" {
			return MoveFile(c, fileProfil)
		}else {
			errUnlink := UnlinkProfile(res.Profil)
			if errUnlink != nil {
				return nil
			}else {
				return MoveFile(c, fileProfil)
			}
		}

	}
}

func MoveFile(c *gin.Context, fileProfil *multipart.FileHeader) *multipart.FileHeader {
	pathProfil := fmt.Sprintf("upload/images/profil/%s", fileProfil.Filename)
	erProf := c.SaveUploadedFile(fileProfil, pathProfil)
	if erProf != nil {
		return nil
	}else {
		return fileProfil
	}
}

func UnlinkProfile(image string) error {
	_, err := os.Stat("./upload/images/profil/" + image)
	if os.IsNotExist(err) {
		return err
	}
	err = os.Remove("./upload/images/profil/" + image)
	if err != nil {
		return err
	}
	return nil
}
