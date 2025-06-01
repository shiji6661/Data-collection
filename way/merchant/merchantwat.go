package merchant

// todo:根据商家id查询
import (
	"errors"
	"models/model_merchant/model_mysql"
)

func GetMerchantIdInfoCommon(id int64) (mer *model_mysql.Merchant, err error) {
	info, err := model_mysql.GetMerchantIdInfo(id)
	if info == nil {
		return nil, errors.New("商家信息查询失败")
	}
	return info, nil
}
