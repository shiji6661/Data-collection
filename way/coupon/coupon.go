package coupon

import "models/model_product/model_mysql"

// 查询优惠卷
func CommonFindCouponById(cId int64) (c *model_mysql.Coupon, err error) {
	c = &model_mysql.Coupon{}
	err = c.FindCouponById(cId)
	if err != nil {
		return nil, err
	}
	return c, nil
}
