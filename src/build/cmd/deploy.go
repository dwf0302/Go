package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CommandDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "项目部署",
	Long:  `请提供部署的（.sh文件）名称,进行部署`,
	Run:   DeployProject,
}

func init() {
	flag := CommandDeploy.Flags()
	flag.StringP("filename", "f", "", "deploy filename")
}

func DeployProject(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("filename")
	fmt.Println(filename)
	//if len(args) == 0 {
	//	fmt.Println("请提供文件名")
	//	return
	//}
	//fmt.Println("文件名称:", args[0])
}
