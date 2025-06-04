package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"product_srv/proto_product/product"
)

const (
	IsPermanent   = 0    // 是否无限张数 0 否 1 是
	CsStatus      = 0    // 状态（0：关闭，1：开启,-1:已失效）
	FullReduction = 1000 // 消费满多少赠送优惠券
)

// TODO：查询当前优惠卷是否存在
func FindCouponByCid(cid int64) (cs *model_mysql.CouponStore, err error) {
	cs = &model_mysql.CouponStore{}
	err = cs.FindCouponStoreByCid(cid)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

// TODO: 商家添加优惠卷
func MerAddCouponStore(in *product.MerAddCouponStoreRequest) (cs *model_mysql.CouponStore, err error) {
	id, err := FindCouponById(in.Cid)
	if err != nil {
		zap.L().Info("查询优惠卷失败")
		return nil, errors.New("查询失败")
	}
	if id.ID == 0 {
		zap.L().Info("优惠卷不存在！请重新选择！")
		return nil, errors.New("优惠卷不存在！请重新选择！")
	}
	// TODO: 判断当前优惠卷是否存在
	cid, err := FindCouponByCid(in.Cid)
	if err != nil {
		zap.L().Info("查询优惠卷失败")
		return nil, errors.New("查询失败")
	}
	if cid.ID != 0 {
		zap.L().Info("当前优惠卷已存在！")
		return nil, errors.New("当前优惠卷已存在！")
	}
	// TODO:添加优惠卷
	if in.StartTime > in.StopTime {
		zap.L().Info("开始时间不能大于结束时间！")
		return nil, errors.New("开始时间不能大于结束时间！")
	}
	if in.Num <= 0 {
		zap.L().Info("优惠卷数量不能小于0！")
		return nil, errors.New("优惠卷数量不能小于0！")
	}

	cs = &model_mysql.CouponStore{
		MerId:         in.MerId,
		Cid:           in.Cid,
		StartTime:     in.StartTime,
		EndTime:       in.StopTime,
		TotalCount:    in.Num,
		RemainCount:   in.Num,
		IsPermanent:   IsPermanent,
		Status:        CsStatus,
		FullReduction: FullReduction,
	}
	err = cs.AddCouponStore()
	if err != nil {
		zap.L().Info("优惠卷添加失败！")
		return nil, errors.New("优惠卷添加失败！")
	}
	return cs, nil
}

// TODO:删除优惠卷
func MerDeleteCoupon(cid int64) (cs *model_mysql.CouponStore, err error) {
	cs = &model_mysql.CouponStore{}
	err = cs.DelCouponStore(cid)
	if err != nil {
		zap.L().Info("优惠卷删除失败！")
		return nil, errors.New("优惠卷删除失败！")
	}
	return cs, nil
}
