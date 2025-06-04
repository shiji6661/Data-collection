package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

func FindLevelInfo() ([]*user.LevelInfo, error) {
	ui := &model_mysql.UserLevel{}
	info, err := ui.FindLevelInfo()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var list []*user.LevelInfo
	for _, i := range info {
		list = append(list, &user.LevelInfo{
			Name:    i.Name,
			Score:   i.Score,
			Img:     i.Img,
			Explain: i.Explain,
		})
	}
	return info, nil
}
