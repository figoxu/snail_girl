package child

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "识别go.txt并删除",
	Long:  `识别go.txt并删除`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please input your work dir")
			return
		}
		for _, v := range args {
			basePath := "/opt" + v
			err := filepath.Walk(basePath,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if strings.HasSuffix(path, ".go.txt") {
						fmt.Println("Delete  ", path, info.Size())
						err = os.Remove(path)
						if err != nil {
							return err
						}
					}
					return nil
				})
			if err != nil {
				fmt.Println(`error at path `, basePath)
				return
			}
		}
		// gen := &writer.ScaffoldDomain{}
		// for _, v := range args {
		// 	gen.GenerateResult(v, "")
		// }
	},
}
