package model

type Soldier struct {
	Id           string //编号
	Rarity       string //稀有度
	UnlockArena  string //解锁阶段
	Cvc          string //客户端版本号
	CombatPoints string //战斗力/战力点
}

func newSoldier(id string, ra string, un string, cvc string, cp string) Soldier {
	return Soldier{
		id, ra, un, cvc, cp,
	}
}

// GetSoldiersByCvc 根据cvc获取所有合法的士兵
func (*Soldier) GetSoldiersByCvc(cvc string, soldiers map[string]Soldier) map[string]Soldier {
	res := make(map[string]Soldier)
	for k, v := range soldiers {
		if v.Cvc == cvc {
			res[k] = newSoldier(v.Id, v.Rarity, v.UnlockArena, v.Cvc, v.CombatPoints)
		}
	}
	return res
}

//GetCombatPointsById 根据士兵id获取战力
func (*Soldier) GetCombatPointsById(id string, soldiers map[string]Soldier) string {
	sol, ok := soldiers[id]
	if !ok {
		return ""
	}
	return sol.CombatPoints
}

//GetRarityById 根据士兵id获取稀有度
func (*Soldier) GetRarityById(id string, soldiers map[string]Soldier) string {
	sol, ok := soldiers[id]
	if !ok {
		return ""
	}
	return sol.Rarity
}

//GetSoldiersByUn 依据解锁阶段分组返回士兵信息
func (*Soldier) GetSoldiersByUn(soldiers map[string]Soldier) map[string][]Soldier {
	res := make(map[string][]Soldier)
	for _, v := range soldiers {
		_, ok := res[v.UnlockArena]
		if ok {
			res[v.UnlockArena] = append(res[v.UnlockArena], v)
		} else {
			res[v.UnlockArena] = make([]Soldier, 0)
		}
	}
	return res
}

//GetSoldiersByRUCv 输入稀有度、当前解锁阶段、cvc。获取该稀有度、cvc合法且已经解锁的所有士兵
func (*Soldier) GetSoldiersByRUCv(ra string, un string, cv string, soldiers map[string]Soldier) map[string]Soldier {
	sols := map[string]Soldier{}
	if un == "" {
		for k, v := range soldiers {
			if v.Rarity == ra && cv == v.Cvc {
				sols[k] = newSoldier(v.Id, v.Rarity, v.UnlockArena, v.Cvc, v.CombatPoints)
			}
		}
	} else {
		for k, v := range soldiers {
			if v.Rarity == ra && Compare(un, v.UnlockArena) && cv == v.Cvc {
				sols[k] = newSoldier(v.Id, v.Rarity, v.UnlockArena, v.Cvc, v.CombatPoints)
			}
		}
	}
	return sols
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
