package main

import (
	"fmt"
	"log"
	"net/http"
	"parse-config/model"
	"parse-config/utils"
	"strings"

	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
)

var serverPort string

// ini配置文件路径
var iniUrl string

//士兵信息json文件的路径
var integralSoldierPath string

//士兵信息新json文件的路径
var partialSoldierPath string

//士兵信息map key：id  value：soldier
var soldiers map[string]model.Soldier

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {
	// 定义命令行参数对应的变量
	jsonPath := flag.StringP("jsonPath", "p", "", "")
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	//接收命令行的json文件路径
	integralSoldierPath = *jsonPath

	partialSoldierPath, soldiers := utils.ReadJson(integralSoldierPath)
	serverPort = utils.ReadIni()
	fmt.Println("新的json地址为", partialSoldierPath)

	r := gin.Default()
	r.GET("/GetSoldiersByCvc", func(c *gin.Context) {
		sol := model.Soldier{}
		cvc := c.DefaultQuery("cvc", "")
		resMap := sol.GetSoldiersByCvc(cvc, soldiers)
		c.JSON(http.StatusOK, resMap)
	})
	r.GET("/GetCombatPointsById", func(c *gin.Context) {
		sol := model.Soldier{}
		id := c.Query("id")
		resMap := sol.GetCombatPointsById(id, soldiers)
		c.JSON(http.StatusOK, resMap)
	})
	r.GET("/GetRarityById", func(c *gin.Context) {
		sol := model.Soldier{}
		id := c.Query("id")
		resMap := sol.GetRarityById(id, soldiers)
		c.JSON(http.StatusOK, resMap)
	})
	r.GET("/GetSoldiersByUn", func(c *gin.Context) {
		sol := model.Soldier{}
		resMap := sol.GetSoldiersByUn(soldiers)
		c.JSON(http.StatusOK, resMap)
	})
	r.GET("/GetSoldiersByRUCv", func(c *gin.Context) {
		sol := model.Soldier{}
		ra := c.Query("Rarity")
		un := c.Query("UnlockArena")
		cv := c.Query("Cvc")
		resMap := sol.GetSoldiersByRUCv(ra, un, cv, soldiers)
		c.JSON(http.StatusOK, resMap)
	})
	go utils.ListenerForJson(&soldiers, integralSoldierPath)
	if err := r.Run(":" + serverPort); err != nil {
		log.Fatal(err)
	}
}
