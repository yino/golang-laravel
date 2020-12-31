package corp

import (
	"github.com/gin-gonic/gin"
	response "go-laravel/app/exception"
)

func FileUploadController(request * gin.Context)  {

	f, err := request.FormFile("img")
	if err != nil {
		response.BaseException(50000,"请上传图片")
		return
	}else{
		request.SaveUploadedFile(f, f.Filename)
		result := make(map[string]string)
		result["url"] = f.Filename
		response.SuccessResponse(result)
	}
}