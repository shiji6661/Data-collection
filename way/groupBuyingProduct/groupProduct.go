package groupBuyingProduct

import "models/model_product/model_mysql"

func ReduceGroupProduct(productId int64, stock int64) (gp *model_mysql.GroupBuyingProduct, err error) {
	gp = &model_mysql.GroupBuyingProduct{}
	err = gp.ReduceStock(productId, stock)
	if err != nil {
		return nil, err
	}
	return gp, nil
}
