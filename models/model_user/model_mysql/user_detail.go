package model_mysql

import "common/global"

// todo:用户详情
type UserDetail struct {
	Id         int64  `gorm:"column:id;type:int UNSIGNED;comment:主键id;primaryKey;not null;" json:"id"` // 主键id
	CreatedAt  string `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt  string `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt  string `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Nickname   string `gorm:"column:nickname;type:varchar(64);comment:用户的昵称;default:NULL;" json:"nickname"`                 // 用户的昵称
	Headimgurl string `gorm:"column:headimgurl;type:varchar(256);comment:用户头像;default:NULL;" json:"headimgurl"`             // 用户头像
	Sex        int64  `gorm:"column:sex;type:tinyint UNSIGNED;comment:用户的性别，值为1时是男性，值为2时是女性，值为0时是未知;default:0;" json:"sex"` // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City       string `gorm:"column:city;type:varchar(64);comment:用户所在城市;default:NULL;" json:"city"`                        // 用户所在城市
	Language   string `gorm:"column:language;type:varchar(64);comment:用户的语言，简体中文为zh_CN;default:NULL;" json:"language"`      // 用户的语言，简体中文为zh_CN
	Province   string `gorm:"column:province;type:varchar(64);comment:用户所在省份;default:NULL;" json:"province"`                // 用户所在省份
	Country    string `gorm:"column:country;type:varchar(64);comment:用户所在国家;default:NULL;" json:"country"`                  // 用户所在国家
	Stair      int64  `gorm:"column:stair;type:int UNSIGNED;comment:一级推荐人;default:NULL;" json:"stair"`                      // 一级推荐人
	Second     int64  `gorm:"column:second;type:int UNSIGNED;comment:二级推荐人;default:NULL;" json:"second"`                    // 二级推荐人
}

func (ud *UserDetail) TableName() string {
	return "user_detail"
}

func (ud *UserDetail) FindUserDetailById(userId int64) (user *UserDetail, err error) {
	err = global.DB.Where("id = ?", userId).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
