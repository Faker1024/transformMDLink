package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/*
拿到地址後，讀取所有Md文檔地址
解析文件內容
獲取文件中圖片鏈接
讀取圖片內容上傳至圖牀，獲得連接
將連接替換爲原來鏈接，保存文件
*/

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "transformMDLink",
	Short: "一个自动将MD文档中的图片上传到图床的命令行工具",
	Long:  `transformMDLink， start！！！`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags().GetString("path"))
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.Flags().String("path", "p", "file path")
}
