package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Configs struct {
	Server struct{
		Port string `yaml:"port"`
		Ip string `yaml:"ip"`
	}
	Database struct{
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
}

func GetConfig() *Configs{
	config := &Configs{}
	absPath, _ := filepath.Abs("/Users/mjea/go/src/gin_mani_center/conf/meta.yaml")
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}
