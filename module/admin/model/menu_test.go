package model

import (
	"testing"
)

func TestMenu1(t *testing.T) {
	var allMenu Menus

	var m = addMenu(0, int64(1))
	allMenu = append(allMenu, m)

	m = addMenu(1, int64(2))
	allMenu = append(allMenu, m)

	m = addMenu(2, int64(3))
	allMenu = append(allMenu, m)

	m = addMenu(3, int64(4))
	allMenu = append(allMenu, m)

	m = addMenu(4, int64(5))
	allMenu = append(allMenu, m)

	var p = &MenuTree{
		ParentId: 0,
		MenuId:   1,
	}
	EachMenuNode(p, allMenu)
	t.Log(p)
}

func TestEachMenu(t *testing.T) {
	var allMenu Menus

	/**
	0 -> 1,2,3
	1 -> 4,5,6
	2 -> 7,8,9
	3 -> 10,11,12
	*/
	//for i := 1; i <= 3; i++ {
	//	var m = addMenu(0, int64(i))
	//	allMenu = append(allMenu, m)
	//}

	//for i := 4; i <= 6; i++ {
	//	var m = addMenu(1, int64(i))
	//	allMenu = append(allMenu, m)
	//}
	//
	//for i := 7; i <= 9; i++ {
	//	var m = addMenu(2, int64(i))
	//	allMenu = append(allMenu, m)
	//}
	//
	//for i := 10; i <= 12; i++ {
	//	var m = addMenu(3, int64(i))
	//	allMenu = append(allMenu, m)
	//}

	//fmt.Print(allMenu)
	var p = &MenuTree{
		ParentId: 0,
		MenuId:   1,
	}
	EachMenuNode(p, allMenu)
	t.Log(p)
}

func addMenu(parent int64, menuId int64) *Menu {
	return &Menu{
		MenuId:   menuId,
		ParentId: parent,
	}
}
