package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configs struct {
	Server struct{
		Port string `yaml:"port"`
		Ip string `yaml:"ip"`
	}
	Database struct{
		Addr string `yaml:"addr"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
	Client struct{
		User string `yaml:"user"`
		Center string `yaml:"center"`
		Engine string `yaml:"engine"`
	}
}

func GetConfig() *Configs{
	config := &Configs{}
	cpath := fmt.Sprintf("%s/src/gin_mani_center/conf/meta.yaml",os.Getenv("GOPATH"))
	content, err := ioutil.ReadFile(cpath)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}
