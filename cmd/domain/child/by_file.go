package child

import (
	"fmt"

	"figoxu.me/snail_girl/pkg/biz/writer"
	"github.com/spf13/cobra"
)

var byFileCmd = &cobra.Command{
	Use:   "by_file",
	Short: "基于文件名去识别domain并处理",
	Long:  `基于文件名去识别domain并处理`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file: ", args)
		gen := &writer.ScaffoldDomain{}
		for _, v := range args {
			gen.GenerateResult(v, "")
		}
	},
}
