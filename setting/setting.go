package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	DbConfig     *DbInfo    `mapstructure:"database"`
	OSSAliConfig *OSSAli    `mapstructure:"ossAli"`
	LogConfig    *LogConfig `mapstructure:"log"`
	OSSQiNiuConfig *OSSQiNiu  `mapstructure:"ossQiNiu"`
}

type DbInfo struct {
	TYPE      string
	USER      string
	PASSWORD  string
	HOST      string
	NAME      string
	PORT      string
}

type OSSAli struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	SufferUrl       string
}

type OSSQiNiu struct {
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiNiuServer string
}

type LogConfig struct {
	Level string `json:"level"`
	Filename string `json:"filename"`
	MaxSize int `json:"maxsize"`
	MaxAge int `json:"max_age"`
	MaxBackups int `json:"max_backups"`
}

func InitConfig() error {
	viper.SetConfigFile("./config/config.yml")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}
