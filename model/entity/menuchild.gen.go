// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameMenuchild = "menuchild"

// Menuchild mapped from table <menuchild>
type Menuchild struct {
	ID       uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Sort     *int64 `gorm:"column:sort;type:bigint" json:"sort"`
	Name     string `gorm:"column:name;type:varchar(256);not null;uniqueIndex:name,priority:1" json:"name"`    // 菜单名
	Ename    string `gorm:"column:ename;type:varchar(256);not null;uniqueIndex:ename,priority:1" json:"ename"` // 英文名
	Logo     string `gorm:"column:logo;type:longtext;not null" json:"logo"`                                    // 图标名
	Link     string `gorm:"column:link;type:varchar(256);not null;uniqueIndex:link,priority:1" json:"link"`    // 路由名
	ParentID uint64 `gorm:"column:parent_id;type:bigint unsigned;not null" json:"parent_id"`                   // 父级id
}

// TableName Menuchild's table name
func (*Menuchild) TableName() string {
	return TableNameMenuchild
}