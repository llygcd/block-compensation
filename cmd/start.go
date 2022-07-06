package cmd

import (
	"github.com/llygcd/block-compensation/config"
	app "github.com/llygcd/block-compensation/internal"
	"github.com/llygcd/block-compensation/internal/global"
	"github.com/llygcd/block-compensation/utils/constant"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
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
	testCmd.Flags().StringVarP(&localConfig, "Config", "c", "", "conf path: /opt/cfg.toml")
}

func test() {
	configPath := "/block-compensation/config"
	if v, ok := os.LookupEnv(constant.EnvNameConfigFilePath); ok {
		configPath = v
	}
	data, err := ioutil.ReadFile(configPath)
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
