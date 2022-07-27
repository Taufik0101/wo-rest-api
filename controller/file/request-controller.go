package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func SaveImageRequest(c *gin.Context) (ktp string, selfie string) {
	fileKtp, errKTP := c.FormFile("file_ktp")
	fileSelfie, errSelfie := c.FormFile("file_selfie")
	if errKTP != nil || fileKtp.Size == 0 || errSelfie != nil || fileSelfie.Size == 0 {
		return "", ""
	}else {
		pathKTP := fmt.Sprintf("upload/images/ktp/%s", fileKtp.Filename)
		pathSelfie := fmt.Sprintf("upload/images/selfie/%s", fileSelfie.Filename)
		errKtp := c.SaveUploadedFile(fileKtp, pathKTP)
		errselfie := c.SaveUploadedFile(fileSelfie, pathSelfie)
		if errKtp != nil || errselfie != nil {
			return "", ""
		}else {
			return fileKtp.Filename, fileSelfie.Filename
		}
	}
}

func RandomPassword(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
