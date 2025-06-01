package coupon

import "models/model_product/model_mysql"

// 查询前台商家是否有优惠卷
func FindCouponById(csId int64) (cs *model_mysql.CouponStore, err error) {
	cs = &model_mysql.CouponStore{}
	err = cs.FindCouponStoreByCid(csId)
	if err != nil {
		return nil, err
	}
	return cs, nil
}
