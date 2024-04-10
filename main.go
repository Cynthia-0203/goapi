package main

import (
	"fmt"
	"os"

	"github.com/Cynthia/goapi/config"
	"github.com/Cynthia/goapi/pkg/cmd"
	"github.com/Cynthia/goapi/pkg/console"
	"github.com/Cynthia/goapi/pkg/database/databaseload"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/redis"
	"github.com/Cynthia/goapi/pkg/validators"
	// "github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func main(){
	
	config.InitConfig()
	
	
	
	// r:=gin.New()
	// route.InitRoute(r)
	
	// logger.InitLogger()
	// databaseload.InitDB()
	// redis.InitRedis()
	// validators.InitRules()
	// gin.SetMode(gin.ReleaseMode)
	
     
	
	// err:=r.Run(":"+viper.GetString("app.port"))

	
	// if err!=nil{
	// 	fmt.Printf("fail to start web service...")
	// }

	 // 应用的主入口，默认调用 cmd.CmdServe 命令
	 var rootCmd = &cobra.Command{
        Use:   "Goapi",
        Short: "A simple forum project",
        Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

        // rootCmd 的所有子命令都会执行以下代码
        PersistentPreRun: func(command *cobra.Command, args []string) {

            // 配置初始化，依赖命令行 --env 参数
            config.InitConfig()

           
			logger.InitLogger()
			databaseload.InitDB()
			redis.InitRedis()
			validators.InitRules()
        },
    }

    // 注册子命令
    rootCmd.AddCommand(
        cmd.CmdServe,
    )

    // 配置默认运行 Web 服务
    cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

    // // 注册全局参数，--env
    // cmd.RegisterGlobalFlags(rootCmd)

    // 执行主命令
    if err := rootCmd.Execute(); err != nil {
        console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
    }
}