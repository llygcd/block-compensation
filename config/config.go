package config

import (
	"bytes"
	"github.com/spf13/viper"
)

type (
	Config struct {
		DataBase DataBaseConf `mapstructure:"database"`
		Server   ServerConf   `mapstructure:"server"`
	}

	DataBaseConf struct {
		Uri    string `mapstructure:"uri"`
		DbName string `mapstructure:"db_name"`
	}

	ServerConf struct {
		App                     string `mapstructure:"app"`
		NodeUrls                string `mapstructure:"node_urls"`
		GrpcUrls                string `mapstructure:"grpc_urls"`
		ChainId                 string `mapstructure:"chain_id"`
		WorkerNumCreateTask     int    `mapstructure:"worker_num_create_task"`
		WorkerNumExecuteTask    int    `mapstructure:"worker_num_execute_task"`
		ThreadNumParseTx        int    `mapstructure:"thread_num_parse_tx"`
		WorkerMaxSleepTime      int    `mapstructure:"worker_max_sleep_time"`
		BlockNumPerWorkerHandle int    `mapstructure:"block_num_per_worker_handle"`

		SleepTimeCreateTaskWorker int `mapstructure:"sleep_time_create_task_worker"`
		MaxConnectionNum          int `mapstructure:"max_connection_num"`
		InitConnectionNum         int `mapstructure:"init_connection_num"`
		ChainBlockInterval        int `mapstructure:"chain_block_interval"`
		BehindBlockNum            int `mapstructure:"behind_block_num"`

		Bech32AccPrefix   string `mapstructure:"bech32_acc_prefix"`
		PromethousPort    string `mapstructure:"promethous_port"`
		OnlySupportModule string `mapstructure:"only_support_module"`
		IsJsonRpcProtocol bool   `mapstructure:"is_json_rpc_protocol"`

		InsertBatchLimit int `mapstructure:"insert_batch_limit"`
	}
)

func ReadConfig(data []byte) (*Config, error) {
	v := viper.New()
	v.SetConfigType("toml")
	reader := bytes.NewReader(data)
	err := v.ReadConfig(reader)
	if err != nil {
		return nil, err
	}
	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
