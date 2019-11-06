package model

type Role struct {
	RoleId   int64  `gorm:"column:role_id;primary_key;auto_increment;comment:'角色id'"`
	RoleName string `gorm:"column:role_name;comment:'角色名称'"`
	MenuList string `gorm:"column:menu_list;comment:'角色菜单id.逗号分隔menu_id'"`
	Remark   string `gorm:"column:remark;comment:'备注'"`
}

func (*Role) TableName() string {
	return "sys_role"
}
