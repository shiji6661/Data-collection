package model_mysql

// todo:商品分类表
type ProductSort struct {
	Id        int64  `gorm:"column:id;type:mediumint;comment:商品分类表ID;primaryKey;not null;" json:"id"`                         // 商品分类表ID
	Pid       int64  `gorm:"column:pid;type:mediumint;comment:父id;default:NULL;" json:"pid"`                                  // 父id
	CateName  string `gorm:"column:cate_name;type:varchar(100);comment:分类名称;default:NULL;" json:"cate_name"`                  // 分类名称
	Sort      int64  `gorm:"column:sort;type:mediumint;comment:排序;default:NULL;" json:"sort"`                                 // 排序
	Pic       string `gorm:"column:pic;type:varchar(128);comment:图标;" json:"pic"`                                             // 图标
	IsShow    int64  `gorm:"column:is_show;type:tinyint(1);comment:是否推荐;default:1;" json:"is_show"`                           // 是否推荐
	CreatedAt string `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
}

func (ps *ProductSort) TableName() string {
	return "product_sort"
}
