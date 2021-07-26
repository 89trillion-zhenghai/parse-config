package handler

import (
	"parse-config/internal/model"
	"parse-config/util"
)

// GetSoldiersByCvc 根据cvc获取所有合法的士兵
func GetSoldiersByCvc(cvc string) []model.Soldier {
	var res []model.Soldier
	for _, v := range util.Soldiers {
		if v.Cvc == cvc {
			res = append(res, v)
		}
	}
	return res
}

//GetCombatPointsById 根据士兵id获取战力
func GetCombatPointsById(id string) string {
	sol, ok := util.Soldiers[id]
	if !ok {
		return ""
	}
	return sol.CombatPoints
}

//GetRarityById 根据士兵id获取稀有度
func GetRarityById(id string) string {
	sol, ok := util.Soldiers[id]
	if !ok {
		return ""
	}
	return sol.Rarity
}

//GetSoldiersByUn 根据解锁阶段分组返回士兵信息
func GetSoldiersByUn(un string) []model.Soldier {
	res := make([]model.Soldier, 0)
	for _, v := range util.Soldiers {
		if v.UnlockArena == un {
			res = append(res, v)
		}
	}
	return res
}

//GetSoldiersByRUCv 输入稀有度、当前解锁阶段、cvc。获取该稀有度、cvc合法且已经解锁的所有士兵
func GetSoldiersByRUCv(ra string, un string, cvc string) []model.Soldier {
	res := make([]model.Soldier, 0)
	if un == "" {
		for _, v := range util.Soldiers {
			if v.Rarity == ra && cvc == v.Cvc {
				res = append(res, v)
			}
		}
	} else {
		for _, v := range util.Soldiers {
			if v.Rarity == ra && Compare(un, v.UnlockArena) && cvc == v.Cvc {
				res = append(res, v)
			}
		}
	}
	return res
}

// Compare 数值字符串比大小,前者大返回true，反之返回false
func Compare(s1 string, s2 string) bool {
	if s2 == "" {
		return false
	}
	//先比较位数再比较大小
	if len(s1) > len(s2) {
		return true
	} else if len(s1) == len(s2) {
		return s1 >= s2
	} else {
		return false
	}
}
