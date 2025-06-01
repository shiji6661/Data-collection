package order

// todo:根据订单id查询
import (
	"errors"
	"models/model_order/model_mysql"
)

func FindByOrderId(id int64) (ord *model_mysql.Order, err error) {
	info, err := model_mysql.FindByOrderId(id)
	if info == nil {
		return nil, errors.New("商家信息查询失败")
	}
	return info, nil
}
