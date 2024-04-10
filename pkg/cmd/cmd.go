package cmd

import (
	"os"

	"github.com/Cynthia/goapi/pkg"
	"github.com/spf13/cobra"
)

var Env string

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {
    rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing will use .env.testing file")
}

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
    cmd, _, err := rootCmd.Find(os.Args[1:])
    firstArg := pkg.FirstElement(os.Args[1:])
    if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
        args := append([]string{subCmd.Use}, os.Args[1:]...)
        rootCmd.SetArgs(args)
    }
}