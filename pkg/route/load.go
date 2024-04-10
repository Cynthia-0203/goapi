package route

import (
	"net/http"
	"strings"

	"github.com/Cynthia/goapi/pkg/middlewares"
	"github.com/Cynthia/goapi/routes"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine){
	initMiddleware(r)
	routes.RegisterAPIRoutes(r)
	setup404(r)
}

func initMiddleware(r *gin.Engine){
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404(r *gin.Engine){
	r.NoRoute(func(c *gin.Context){
		acceptString:=c.Request.Header.Get("Accept")
		if strings.Contains(acceptString,"text/html"){
			c.String(http.StatusNotFound,"404")
		}else{
			c.JSON(http.StatusNotFound,gin.H{
				"error_code":404,
				"error_message":"route is not defined...",
			})
		}
	})
}