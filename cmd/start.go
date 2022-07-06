package cmd

import (
	"github.com/llygcd/block-compensation/config"
	app "github.com/llygcd/block-compensation/internal"
	"github.com/llygcd/block-compensation/internal/global"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var (
	localConfig string
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Start Block Compensation",
		Run: func(cmd *cobra.Command, args []string) {
			test()
		},
	}
	testCmd = &cobra.Command{
		Use:   "start",
		Short: "Start  Block Compensation",
		Run: func(cmd *cobra.Command, args []string) {
			test()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&localConfig, "Config", "c", "", "conf path: /home/lly/go/src/github.com/llygcd/block-compensation/config/cfg.toml")
}

func test() {
	data, err := ioutil.ReadFile("/home/lly/go/src/github.com/llygcd/block-compensation/config/cfg.toml")
	if err != nil {
		panic(err)
	}
	config, err := config.ReadConfig(data)
	if err != nil {
		panic(err)
	}
	run(config)
}

func run(cfg *config.Config) {
	global.Config = cfg
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetLevel(log.DebugLevel)
	app.Serve(cfg)
}
