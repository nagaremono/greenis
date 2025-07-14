package internal

import (
	"flag"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rdbCfg *RDBConfig

type RDBConfig struct {
	Dir      string
	Filename string
	mu       sync.RWMutex
}

func (r *RDBConfig) SetDir(d string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Dir = d
}

func (r *RDBConfig) SetFile(f string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.Filename = f
}

func (r *RDBConfig) GetDir() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	return viper.GetString("dir")
}

func (r *RDBConfig) GetFile() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	return viper.GetString("dbfilename")
}

func GetRDBConfig() *RDBConfig {
	if rdbCfg != nil {
		return rdbCfg
	}

	viper.SetDefault("dir", "/tmp/greenis-data")
	viper.SetDefault("dbfilename", "rdbfile")

	flag.String("dir", "", "directory for rdb config file")
	flag.String("dbfilename", "", "rdb config file name")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	rdbCfg = &RDBConfig{}

	return rdbCfg
}
