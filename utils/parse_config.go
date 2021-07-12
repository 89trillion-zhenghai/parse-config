package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"parse-config/model"

	"github.com/fsnotify/fsnotify"

	"gopkg.in/ini.v1"
)

//ReadJson 读取json文件,提取有用的信息并且保存到configNew.army.model.json配置文件中
func ReadJson(fn string) (string, map[string]model.Soldier) {
	sd := make(map[string]model.Soldier)
	bytes, err := ioutil.ReadFile(fn)
	getError("file not find or other error!", err)
	err = json.Unmarshal(bytes, &sd)
	getError("error, please check!", err)
	marshal, err := json.Marshal(sd)
	getError("error, please check!", err)
	newJson, err := os.Create("configNew.army.model.json")
	defer newJson.Close()
	newJson.Write(marshal)
	return newJson.Name(), sd
}

//ReadIni 读取app.ini配置文件，遍历所有分区，找到HttpPort,输出http端口号
func ReadIni() string {
	conf, err := ini.Load("./resource/app.ini")
	getError("file not find or other error!", err)
	return conf.Section("server").Key("HttpPort").String()
}

//错误处理
func getError(msg string, err error) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

//ListenerForJson 监听json文件，发生update操作则对Soldier和new.json更新
func ListenerForJson(sod *map[string]model.Soldier, fn string) {
	watch, err := fsnotify.NewWatcher()
	getError("error!", err)
	defer watch.Close()
	err = watch.Add(fn)
	getError("error!", err)
	func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Write == fsnotify.Write {
						//更新json文件  步骤： 根据最新的json提取map ，新建一个config.army.model_sort.json，如果前两步都正确，则将旧文件替换掉
						fmt.Println("配置文件更新......")
						UpdateMapAndJson(sod, fn)
						fmt.Println("配置文件更新成功")
					}
				}
			case err := <-watch.Errors:
				{
					fmt.Println(err)
					return
				}
			}
		}
	}()
	//循环
	select {}
}

func UpdateMapAndJson(sols *map[string]model.Soldier, fn string) {
	sd := make(map[string]model.Soldier)
	bytes, err := ioutil.ReadFile(fn)
	getError("file not find or other error!", err)
	err = json.Unmarshal(bytes, &sd)
	getError("error, please check!", err)
	marshal, err := json.Marshal(sd)
	getError("error, please check!", err)
	newJson, err := os.Create("configNew.army.model.json")
	getError("create file error!", err)
	defer newJson.Close()
	newJson.Write(marshal)
	*sols = sd
}
