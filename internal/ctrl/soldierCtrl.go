package ctrl

import (
	"parse-config/internal/globalError"
	"parse-config/internal/handler"
	"parse-config/internal/verify"

	"github.com/gin-gonic/gin"
)

// GetSoldiersByCvc 根据cvc获取所有合法的士兵
func GetSoldiersByCvc(c *gin.Context) (interface{}, error) {
	cvc := c.Query("cvc")
	err := verify.ParamIsEmpty(cvc)
	if err != nil {
		return nil, err
	}
	err = verify.ParamIsDigit(cvc)
	if err != nil {
		return nil, err
	}
	res := handler.GetSoldiersByCvc(cvc)
	return res, nil
}

//GetCombatPointsById 根据士兵id获取战力
func GetCombatPointsById(c *gin.Context) (interface{}, error) {
	id := c.Query("id")
	err := verify.ParamIsEmpty(id)
	if err != nil {
		return nil, err
	}
	err = verify.ParamIsDigit(id)
	if err != nil {
		return nil, err
	}
	combatPoints := handler.GetCombatPointsById(id)
	if len(combatPoints) == 0 {
		return nil, globalError.ResultExpression("士兵不存在")
	}
	return combatPoints, nil
}

//GetRarityById 根据士兵id获取稀有度
func GetRarityById(c *gin.Context) (interface{}, error) {
	id := c.Query("id")
	err := verify.ParamIsEmpty(id)
	if err != nil {
		return nil, err
	}
	err = verify.ParamIsDigit(id)
	if err != nil {
		return nil, err
	}
	rarity := handler.GetRarityById(id)
	if len(rarity) == 0 {
		return nil, globalError.ResultExpression("士兵不存在")
	}
	return rarity, nil
}

//GetSoldiersByUn 依据解锁阶段返回士兵信息
func GetSoldiersByUn(c *gin.Context) (interface{}, error) {
	un := c.Query("un")
	err := verify.ParamIsEmpty(un)
	if err != nil {
		return nil, err
	}
	err = verify.ParamIsDigit(un)
	if err != nil {
		return nil, err
	}
	soldiers := handler.GetSoldiersByUn(un)
	return soldiers, nil
}

//GetSoldiersByRUCv 输入稀有度、当前解锁阶段、cvc。获取该稀有度、cvc合法且已经解锁的所有士兵
func GetSoldiersByRUCv(c *gin.Context) (interface{}, error) {
	ra := c.Query("Rarity")
	un := c.Query("UnlockArena")
	cvc := c.Query("Cvc")
	err := verify.ParamIsEmpty(ra, un, cvc)
	if err != nil {
		return nil, err
	}
	err = verify.ParamIsDigit(ra, un, cvc)
	if err != nil {
		return nil, err
	}
	soldiers := handler.GetSoldiersByRUCv(ra, un, cvc)
	return soldiers, nil
}
