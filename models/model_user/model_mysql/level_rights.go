package model_mysql

import (
	"common/global"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

type LevelRights struct {
	Id          int64 `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	UserLevelId int64 `gorm:"column:user_level_id;type:int;comment:用户等级id;not null;" json:"user_level_id"` // 用户等级id
	RightsId    int64 `gorm:"column:rights_id;type:int;comment:权益id;not null;" json:"rights_id"`           // 权益id
}

func (LevelRights) TableName() string {
	return "level_rights"
}

func (r *LevelRights) FindUserLevelRights(lid int64) (ri []*user.RightsInfo, err error) {
	err = global.DB.Debug().Raw("SELECT  rights.rights_name,  rights.user_level_id,  rights.img,  rights.`explain` FROM  level_rights  JOIN rights ON level_rights.id = rights.user_level_id WHERE  level_rights.user_level_id = ?", lid).Scan(&ri).Error
	if err != nil {
		return nil, err
	}
	return ri, nil
}

func (r *LevelRights) FindULRById(lid int64) (lr *LevelRights, err error) {
	err = global.DB.Debug().Where("user_level_id = ?", lid).Find(&lr).Error
	if err != nil {
		return nil, err
	}
	return lr, nil
}
