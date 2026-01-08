package do

type UserDO struct {
	BaseModel
	// 关键修改：添加 type:varchar(255) 或者 size:255
	// 另外推荐使用 uniqueIndex 替代 index:...,unique
	Username string `vd:"len($)>5" gorm:"type:varchar(255);uniqueIndex"`

	Nickname string `gorm:"type:varchar(255)"` // 建议 Nickname 也加上长度限制
	Password string `gorm:"type:varchar(255)"` // 建议 Password 也加上
	Role     string `gorm:"type:varchar(50)"`
	CanUse   bool
}
