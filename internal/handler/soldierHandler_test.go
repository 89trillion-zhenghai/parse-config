package handler

import (
	"testing"
)

func TestGetSoldiersByCvc(t *testing.T) {
	soldier := GetSoldiersByUn("0")
	t.Log(soldier)
}

func TestGetCombatPointsById(t *testing.T) {
	combatPoints := GetCombatPointsById("10101")
	t.Log(combatPoints)
}

func TestGetRarityById(t *testing.T) {
	rarity := GetRarityById("10101")
	t.Log(rarity)
}

func TestGetSoldiersByUn(t *testing.T) {
	soldiers := GetSoldiersByUn("2")
	t.Log(soldiers)
}

func TestGetSoldiersByRUCv(t *testing.T) {
	soldiers := GetSoldiersByRUCv("2", "3", "1000")
	t.Log(soldiers)
}
