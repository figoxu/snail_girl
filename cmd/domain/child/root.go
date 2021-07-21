package child

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snailgirl",
	Short: "田螺姑娘",
	Long:  `田螺姑娘`,
}

func Execute() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(byDirCmd)
	rootCmd.AddCommand(byFileCmd)
	rootCmd.AddCommand(dockerByFileCmd)
	rootCmd.AddCommand(cleanCmd)
}
