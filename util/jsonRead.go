package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"parse-config/internal/model"
	"strings"

	flag "github.com/spf13/pflag"
)

//Soldiers 士兵信息map key：id  value：soldier
var Soldiers map[string]model.Soldier

func Init() {
	// 定义命令行参数对应的变量
	jsonPath := flag.StringP("jsonPath", "p", "", "")
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	fs := *jsonPath
	Soldiers = ReadJson(fs)
}

//ReadJson 读取json文件,提取有用的信息并且保存到configNew.army.model.json配置文件中
func ReadJson(fs string) map[string]model.Soldier {
	var err error
	soldiers := make(map[string]model.Soldier)
	fmt.Println(fs)
	file, err := os.Open(fs)
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	err = json.Unmarshal(bytes, &soldiers)
	marshal, err := json.Marshal(soldiers)
	newJson, err := os.Create("../conf/soldier.json")
	defer newJson.Close()
	newJson.Write(marshal)
	if err != nil {
		fmt.Println(err.Error())
	}
	return soldiers
}

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}
