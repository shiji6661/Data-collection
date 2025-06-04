package dao_mysql

import (
	"Data-collection/way/coupon"
	"errors"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"time"
	"user_srv/proto_user/user"

	user2 "Data-collection/way/user"
)

// TODO:查询用户是否已经领取过该优惠卷
func IsUserReceiveCou(uid, csId int64) (cu *model_mysql.CouponUserUse, err error) {
	cu = &model_mysql.CouponUserUse{}
	err = cu.FindCouponByUidAndCid(uid, csId)
	if err != nil {
		return nil, err
	}
	return cu, nil
}

// todo:用户领取优惠卷
func UserReceiveCoupon(in *user.UserReceiveCouponRequest) (cu *model_mysql.CouponUserUse, err error) {
	csId, err := coupon.FindCouponById(in.CSId)
	if err != nil {
		return nil, err
	}
	if csId.ID == 0 {
		zap.L().Info("优惠卷不存在！请联系商家！")
		return nil, errors.New("优惠卷不存在！请联系商家！")
	}

	cId, err := coupon.CommonFindCouponById(in.CSId)
	if err != nil {
		return nil, err
	}
	if cId.ID == 0 {
		zap.L().Info("优惠卷不存在！请联系商家！")
		return nil, errors.New("优惠卷不存在！请联系商家！")
	}

	// todo:判断用户是否已经领取过该优惠卷
	cou, err := IsUserReceiveCou(in.Userid, in.CSId)
	if err != nil {
		return nil, err
	}
	if cou.ID != 0 {
		zap.L().Info("用户已经领取过该优惠卷！")
		return nil, errors.New("用户已经领取过该优惠卷！")
	}

	// todo:判断用户积分是否足够
	users, err := user2.CommonGetUserIdInfo(in.Userid)
	if err != nil {
		return nil, err
	}
	if users.ID == 0 {
		zap.L().Info("不存在该用户！请重新登录！")
		return nil, errors.New("不存在该用户！请重新登录！")
	}
	if users.UserPoint < cId.Integral {
		zap.L().Info("用户积分不足！")
		return nil, errors.New("用户积分不足！")
	}
	// todo:判断优惠卷是否已经发放完毕
	if csId.RemainCount == 0 {
		zap.L().Info("优惠卷已经发放完毕！")
		return nil, errors.New("优惠卷已经发放完毕！")
	}
	// todo:判断优惠卷时间是否在领取时间之内
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if nowTime < csId.StartTime || nowTime > csId.EndTime {
		zap.L().Info("优惠卷时间不在领取时间之内！")
		return nil, errors.New("优惠卷时间不在领取时间之内！")
	}

	cu = &model_mysql.CouponUserUse{
		Cid:         csId.Cid,
		Uid:         in.Userid,
		SCId:        in.CSId,
		CouponTitle: cId.Title,
		CouponPrice: cId.CouponPrice,
		Status:      0,
		IsFail:      0,
	}
	err = cu.ReceiveCoupon()
	if err != nil {
		zap.L().Info("领取失败！")
		return nil, errors.New("领取失败！")
	}
	return cu, nil
}
