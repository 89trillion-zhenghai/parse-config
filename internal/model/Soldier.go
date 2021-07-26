package model

type Soldier struct {
	Id           string `json:"id"`           //编号
	Rarity       string `json:"Rarity"`       //稀有度
	UnlockArena  string `json:"UnlockArena"`  //解锁阶段
	Cvc          string `json:"Cvc"`          //客户端版本号
	CombatPoints string `json:"CombatPoints"` //战斗力/战力点
}
