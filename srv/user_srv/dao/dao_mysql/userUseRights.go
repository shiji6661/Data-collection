package dao_mysql

import (
	"errors"
	"user_srv/proto_user/user"
)

func GetUserUseRightsList(uid int64) (newList []*user.UseRights, err error) {
	getUseRights, err := U.GetUserUseRights(uid)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var news []*user.UseRights
	for _, i := range getUseRights {
		news = append(news, &user.UseRights{
			UserName:   i.UserName,
			RightsName: i.RightsName,
			Img:        i.Img,
			Explain:    i.Explain,
			CreatedAt:  i.CreatedAt,
		})
	}
	return news, nil
}
