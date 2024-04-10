package response

import (
	"net/http"

	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JSON(c *gin.Context,data interface{}){
	c.JSON(http.StatusOK,data)
}

func Success(c *gin.Context){
	JSON(c,gin.H{
		"success":true,
		"message":"make it!",
	})
}

func Data(c *gin.Context,data interface{}){
	JSON(c,gin.H{
		"success":true,
		"data":data,
	})
}

func Created(c *gin.Context,data interface{}){
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"data":data,
	})
}

func CreatedJSON(c *gin.Context,data interface{}){
	c.JSON(http.StatusCreated,data)
}

func defaultMessage(defaultMsg string,msg ...string)(message string){
	if len(msg)>0{
		message=msg[0]
	}else{
		message =defaultMsg
	}
	return
}


func Abort404(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
		"message":defaultMessage("data is not exist...",msg...),
	})
}

func Abort500(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
		"message":defaultMessage("the internal server arise error...",msg...),
	})
}

func BadRequest(c *gin.Context,err error,msg ...string){
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
		"message":defaultMessage("fail to parse request ,use multipart and json...",msg...),
		"error":err.Error(),
	})
}

func Error(c *gin.Context,err error,msg ...string){
	logger.LogIf(err)
	if err==gorm.ErrRecordNotFound{
		Abort404(c)
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		"message":defaultMessage("failed to handle request...",msg...),
		"error":err.Error(),
	})
}

func ValidationError(c *gin.Context,errors map[string][]string){
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		"message":"failed to accept request...",
		"errors":errors,
	})
}

func Unauthorized(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
		"message":defaultMessage("failed to parse request...",msg...),
	})
}