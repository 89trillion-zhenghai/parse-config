![未命名文件](https://user-images.githubusercontent.com/86946999/125233602-83326800-e311-11eb-9d89-da672fe2e4b5.jpg)
## 项目简介

1、整理士兵配置文件格式且只保留有用的数据，利用gin开发服务响应用户请求并返回数据。



## 快速上手

1、本项目由Go语言开发，需要配置Go开发环境

2、进入项目根目录，通过命令行命令启动项目

​		1、go build main.go		---编译

​		2、./main -p=./resource/config.army.model.json		---通过命令行输入json路径并启动项目

## 功能介绍

1、读取app.ini配置文件，监听配置中的http端口号。

2、读取命令行中保存士兵信息的json文件路径，对其格式化并取得有效数据。将其保存在新的json文件中。

3、项目启动解析新的json文件，获取士兵信息保存在全局变量soldiers中。

4、项目启动后，对以下请求进行相应并且返回需要的数据信息。

​	1、输入稀有度、当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵

​	2、输入士兵id获取稀有度

​	3、输入士兵id获取战力

​	4、输入cvc获取所有合法的士兵

​	5、获取每个阶段解锁相应士兵的json数据

5、项目运行后，当原士兵配置文件改变后，会重新读取并更新新的json文件内容和存储士兵的全局变量soldiers。

## 代码实现

​	1、士兵信息类

```go
type Soldier struct {
   Id           string //编号
   Rarity       string //稀有度
   UnlockArena  string //解锁阶段
   Cvc          string //客户端版本号
   CombatPoints string //战斗力/战力点
}
func newSoldier(id string, ra string, un string, cvc string, cp string) Soldier {...}
```

​	2、五个处理请求的方法

```go
// GetSoldiersByCvc 根据cvc获取所有合法的士兵
func (*Soldier) GetSoldiersByCvc(cvc string, soldiers map[string]Soldier) map[string]Soldier {...}

//GetCombatPointsById 根据士兵id获取战力
func (*Soldier) GetCombatPointsById(id string, soldiers map[string]Soldier) string {...}

//GetRarityById 根据士兵id获取稀有度
func (*Soldier) GetRarityById(id string, soldiers map[string]Soldier) string {...}

//GetSoldiersByUn 依据解锁阶段分组返回士兵信息
func (*Soldier) GetSoldiersByUn(soldiers map[string]Soldier) map[string][]Soldier {...}

//GetSoldiersByRUCv 输入稀有度、当前解锁阶段、cvc。获取该稀有度、cvc合法且已经解锁的所有士兵
func (*Soldier) GetSoldiersByRUCv(ra string, un string, cv string, soldiers map[string]Soldier) map[string]Soldier {...}
```

​	3、解析和更新配置文件

```go
//ReadJson 读取json文件,提取有用的信息并且保存到configNew.army.model.json配置文件中
func ReadJson(fn string) (string, map[string]model.Soldier) {...}

//ReadIni 读取app.ini配置文件，遍历所有分区，找到HttpPort,返回http端口号
func ReadIni() string {...}

//错误处理
func getError(msg string, err error) {...}

//ListenerForJson 监听json文件，发生update操作则对Soldier和new.json更新
func ListenerForJson(sod *map[string]model.Soldier, fn string) {...}

//UpdateMapAndJson 更新配置文件和Soldiers
func UpdateMapAndJson(sols *map[string]model.Soldier, fn string) {...}
```
