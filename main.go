package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"go_base_demo/test"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	fmt.Println("test")

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

	test.Test()
	test.KlogAndError()

}

type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
}

var db *gorm.DB

func (c *conf) getConf() *conf {

	confYaml, err := ioutil.ReadFile("C:\\xuexi\\github\\go_base_demo\\yaml\\conf.yaml")

	if err != nil {
		fmt.Println(err.Error())
	}

	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(confYaml, c)

	if err != nil {
		panic(err.Error())
	}

	/*db, err = gorm.Open("mysql", "<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}*/
	return c
}
