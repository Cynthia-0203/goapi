package cmd

import (
	"github.com/Cynthia/goapi/pkg/console"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CmdServe=&cobra.Command{
	Use: "serve",
	Short: "Start web server",
	Run: runWeb,
	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command,args []string){
	gin.SetMode(gin.ReleaseMode)

	router:=gin.New()
	route.InitRoute(router)
	err:=router.Run(":"+viper.GetString("app.port"))
	if err!=nil{
		logger.ErrorString("CMD","serve",err.Error())
		console.Exit("Unable to start server,error:"+err.Error())
	}
}