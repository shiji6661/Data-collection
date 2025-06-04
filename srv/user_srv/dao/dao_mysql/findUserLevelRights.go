package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

func FindUserLevelRights(lid int64) (ri []*user.RightsInfo, err error) {
	LR := &model_mysql.LevelRights{}

	getULR, err := LR.FindULRById(lid)

	if getULR.UserLevelId == 0 {
		return nil, errors.New("等级不存在")
	}

	getLR, err := LR.FindUserLevelRights(lid)

	var newri []*user.RightsInfo
	for _, i := range getLR {
		newri = append(newri, &user.RightsInfo{
			RightsName:  i.RightsName,
			UserLevelId: i.UserLevelId,
			Img:         i.Img,
			Explain:     i.Explain,
		})
	}

	return newri, nil
}
