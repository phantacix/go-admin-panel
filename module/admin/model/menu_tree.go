package model

type MenuTree struct {
	MenuId   int64
	ParentId int64
	Text     string
	Url      string
	Icon     string
	OrderNum int

	State     map[string]bool
	Checked   bool
	Children  []*MenuTree
	HasParent bool
}

func (m *MenuTree) HasChild() bool {
	return len(m.Children) > 0
}

type MenuTrees []*MenuTree
