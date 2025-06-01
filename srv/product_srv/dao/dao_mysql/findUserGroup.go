package dao_mysql

import "models/model_product/model_mysql"

// 通过拼团商品id查询拼团信息
func FindUserGroupByCid(cid int64) (gb *model_mysql.GroupBuying, err error) {
	gb = &model_mysql.GroupBuying{}
	err = gb.FindGroupBuyingById(cid)
	if err != nil {
		return nil, err
	}
	return gb, nil
}
