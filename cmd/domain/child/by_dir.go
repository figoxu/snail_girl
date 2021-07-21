package child

import (
	"fmt"
	"io/ioutil"

	"figoxu.me/snail_girl/pkg/biz/sniffer"
	"figoxu.me/snail_girl/pkg/ut"
	"github.com/spf13/cobra"
)

var byDirCmd = &cobra.Command{
	Use:   "by_dir",
	Short: "基于目录去识别domain并处理",
	Long:  `基于目录去识别domain并处理`,
	Run: func(cmd *cobra.Command, args []string) {
		gps := &sniffer.GoPkgSniffer{}
		for _, v := range args {
			files, err := ioutil.ReadDir(v)
			ut.Chk(err)
			for _, file := range files {
				filePath := v + "/" + file.Name()
				fmt.Println(filePath)
				v, err := gps.AsPkg(filePath)
				ut.Chk(err)
				fmt.Println("-->   " + v + "    <--")
			}
		}
		// gps := &sniffer.GoPkgSniffer{}
		// v, err := gps.AsPkg("/Users/xujianhui/xxbmm/projects/workspace_go/meishi/m2/pkg/domain/buy_group.go")
		// ut.Chk(err)
		// fmt.Println(v)
	},
}
