package router

import (
	"parse-config/internal/ctrl"
	"parse-config/internal/globalError"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	//根据cvc获取所有合法的士兵
	r.GET("/GetSoldiersByCvc", globalError.ErrorHandler(ctrl.GetSoldiersByCvc))
	//根据士兵id获取战力
	r.GET("/GetCombatPointsById", globalError.ErrorHandler(ctrl.GetCombatPointsById))
	//根据士兵id获取稀有度
	r.GET("/GetRarityById", globalError.ErrorHandler(ctrl.GetRarityById))
	//根据解锁阶段分组返回士兵信息
	r.GET("/GetSoldiersByUn", globalError.ErrorHandler(ctrl.GetSoldiersByUn))
	//根据稀有度、当前解锁阶段、cvc。获取该稀有度、cvc合法且已经解锁的所有士兵
	r.GET("/GetSoldiersByRUCv", globalError.ErrorHandler(ctrl.GetSoldiersByRUCv))
	return r
}
