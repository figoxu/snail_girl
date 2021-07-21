package child

import (
	"fmt"
	"io/ioutil"

	"figoxu.me/snail_girl/pkg/biz/writer"
	"figoxu.me/snail_girl/pkg/ut"
	"github.com/ahmetb/go-linq/v3"
	"github.com/spf13/cobra"
)

var dockerByFileCmd = &cobra.Command{
	Use:   "docker_by_file",
	Short: "基于docker环境,结构体名称去识别domain并处理",
	Long:  `基于docker环境,结构体名称去识别domain并处理`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("请输入 目录和domain的名称")
			return
		}
		fmt.Println("file: ", args)
		basePath := "/opt" + args[0]
		fmt.Println(basePath)

		structName := args[1]
		fmt.Println("structName ", structName)
		files, err := ioutil.ReadDir(basePath)
		ut.Chk(err)
		fu := &ut.FileUt{}
		for _, file := range files {
			fullPath := basePath + "/" + file.Name()
			vs, err := fu.ReadLinesSlice(fullPath)
			ut.Chk(err)
			matchFlag := linq.From(vs).WhereT(func(x string) bool {
				prepareExp := fmt.Sprintf(`.*type\s+%v\s+struct\s+\{.*`, structName)
				parser := ut.Parser{
					PrepareReg: []string{prepareExp},
					ProcessReg: []string{},
				}
				contents := parser.Exe(x)
				return len(contents) > 0
			}).Count() > 0
			if matchFlag {
				fmt.Println(fullPath, "is match ")
				gen := &writer.ScaffoldDomain{}
				gen.GenerateResult(fullPath, structName)
				fmt.Println("All Commands Done !! ")
				return
			}
		}
		fmt.Println(" Sorry ,We Cant Not Found Domain ", structName, " In FilePath ", basePath)
	},
}
