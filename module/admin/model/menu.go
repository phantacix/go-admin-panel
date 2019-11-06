package model

import "github.com/phantacix/go-admin-panel/core/db"

type Menu struct {
	MenuId   int64  `gorm:"column:menu_id;primary_key;auto_increment;comment:'菜单id'"`
	ParentId int64  `gorm:"column:parent_id;comment:'菜单父id'"`
	Name     string `gorm:"column:name;comment:'菜单名称'"`
	Url      string `gorm:"column:url;comment:'菜单url'"`
	AuthFlag string `gorm:"column:auth_flag;comment:'授权标识'"`
	Type     int    `gorm:"column:type;comment:'菜单类型(0:目录,1:菜单,2:按钮)'"`
	Icon     string `gorm:"column:icon;comment:'菜单图标'"`
	OrderNum int    `gorm:"column:order_num;comment:'排序id'"`
}

type Menus []*Menu

var tableName = "sys_menu"

func (m *Menu) ToMenuTree() *MenuTree {
	var t = &MenuTree{
		MenuId:   m.MenuId,
		ParentId: m.ParentId,
		Text:     m.Name,
		Url:      m.Url,
		Icon:     m.Icon,
		OrderNum: m.OrderNum,
	}
	return t
}

func (*Menu) TableName() string {
	return tableName
}

func (m *Menu) Save() {
	db.AdminDB().Save(m)
}

func (m *Menu) Update() {
	db.AdminDB().Update(m)
}

func (m *Menu) Remove() {
	db.AdminDB().Delete(m)
}

func (m *Menu) Insert() {
	db.AdminDB().Create(m)
}

// GetMenuList 获取所有Menu列表
func MenuList() Menus {
	var menus Menus
	db.AdminDB().Find(&menus)
	return menus
}

// GetSubMenuCount 统计menuId下面子Menu总数
func SubMenuCount(menuId int64) int64 {
	var count int64
	db.AdminDB().Table(tableName).Where("parent_id=?", menuId).Count(count)
	return count
}

// GetMenuTreesByRole 根据roleId获取MenuTree
func MenuTreesByRole(roleId int64) MenuTrees {
	var allMenu Menus
	db.AdminDB().Find(allMenu)

	var tree MenuTrees

	//构建所有MenuTree
	if roleId < 1 {
		tree = BuildMenuTree(allMenu)
	}

	//根据roleId已有的Menu构建MenuTree
	return tree
}

// BuildMenuTree 根据Menus构造MenuTree
func BuildMenuTree(allMenu Menus) MenuTrees {
	var menuTrees MenuTrees

	for _, v := range allMenu {
		if v.ParentId == 0 {
			menuTrees = append(menuTrees, v.ToMenuTree())
		}
	}
	for _, v := range menuTrees {
		EachMenuNode(v, allMenu)
	}
	return menuTrees
}

// EachMenuNode 遍历Menu对象，构建父子对象
func EachMenuNode(parent *MenuTree, allMenu Menus) {
	for _, v := range allMenu {
		if parent.MenuId == v.ParentId {
			var m = &MenuTree{
				MenuId:    v.MenuId,
				ParentId:  parent.MenuId,
				Text:      v.Name,
				Url:       v.Url,
				Icon:      v.Icon,
				OrderNum:  v.OrderNum,
				HasParent: true,
			}
			parent.Children = append(parent.Children, m)
			EachMenuNode(m, allMenu)
		}
	}
}
