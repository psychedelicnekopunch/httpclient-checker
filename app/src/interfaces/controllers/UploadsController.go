
package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type UploadsController struct {

}


func NewUploadsController() *UploadsController {
	return &UploadsController{}
}

// ファイルアップロード / 単一のファイル
// https://gin-gonic.com/ja/docs/examples/upload-file/single-file/
//
// ファイルアップロード / 複数のファイル
// https://gin-gonic.com/ja/docs/examples/upload-file/multiple-file/
//
func (controller *UploadsController) Post(c *gin.Context) {

	// form, _ := c.MultipartForm()
	// files := form.File["test"]
	filename := ""
	file, err := c.FormFile("test")
	if err != nil {
		fmt.Print("\n=========\n", err.Error(), "\n=========\n")
	}
	if err == nil {
		filename = file.Filename
	}

	c.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"method": "POST",
		"contentType": c.GetHeader("Content-Type"),
		"filename": filename,
	})
}
