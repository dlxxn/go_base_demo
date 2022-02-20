package conf

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func YamlTest() {
	//读取yaml配置文件
	var c conf
	confYaml := c.getConf()
	//fmt.Println(c.User)
	//将对象转换成json格式
	jsonConf, err := json.Marshal(confYaml)

	if err != nil {
		fmt.Println("error is\t", err.Error())
	}

	fmt.Println("jsomYaml is\t", string(jsonConf))
}

type conf struct {
	Host string `conf:"host"`
	User string `conf:"user"`
	Pwd  string `conf:"pwd"`
}

func (c *conf) getConf() *conf {

	confYaml, err := ioutil.ReadFile("C:\\xuexi\\github\\go_base_demo\\conf\\conf.yaml")

	if err != nil {
		fmt.Println(err.Error())
	}

	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(confYaml, c)

	if err != nil {
		panic(err.Error())
	}

	return c
}
