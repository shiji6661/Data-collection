package trigger

import (
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"product_srv/proto_product/product"
)

// TODO：商品详情
func ProductInfo(c *gin.Context) {
	var data request.ProductInfo
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	ProductInfos, err := handler.ProductInfo(c, &product.ProductInfoRequest{
		StoreName: data.Name,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, ProductInfos)
}

// todo:商品推荐
func ProductRecommend(c *gin.Context) {

	ProductInfos, err := handler.ProductRecommend(c, &product.ProductRecommendRequest{})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, ProductInfos)
}

// todo:商品关键字查询
func ProductKeyword(c *gin.Context) {
	var data request.Keyword
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	ProductInfos, err := handler.ProductKeyword(c, &product.ProductKeywordRequest{
		Keyword: data.Keyword,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, ProductInfos)
}

// todo:商品筛选
func ProductFilter(c *gin.Context) {
	var data request.Filter
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	ProductInfos, err := handler.ProductFilter(c, &product.ProductFilterRequest{
		MinPrice:  data.Min_price,
		MaxPrice:  data.Max_price,
		CateId:    data.Cate_id,
		Sales:     data.Sales,
		IsShow:    data.Is_show,
		IsHot:     data.Is_hot,
		IsBenefit: data.Is_benefit,
		IsBest:    data.Is_best,
		IsNew:     data.Is_new,
		IsSeckill: data.Is_seckill,
		IsBargain: data.Is_bargain,
		Name:      data.Name,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, ProductInfos)
}

// todo:商品分类查询
func ProductCategory(c *gin.Context) {
	var data request.ProductCategory
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	ProductInfos, err := handler.ProductCategory(c, &product.ProductCategoryRequest{
		CateId: data.CateId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, ProductInfos)
}

//// todo:添加购物车
//func AddToCart(c *gin.Context) {
//	var data request.AddToCart
//	if err := c.ShouldBind(&data); err != nil {
//		response.ResponseError(c, err.Error())
//		return
//	}
//
//	id := c.GetUint("userId")
//
//	ProductInfos, err := handler.AddToCart(c, &product.AddToCartRequest{
//		UserId:    int64(id),
//		ProductId: data.ProductId,
//		Num:       data.Num,
//	})
//	if err != nil {
//		response.ResponseError(c, err.Error())
//		return
//	}
//	response.ResponseSuccess(c, ProductInfos.Success)
//
//}

func CreateBargainProduct(c *gin.Context) {
	var data request.CreateBargainProductRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	fmt.Println(222)
	fmt.Println(data)
	bargainProduct, err := handler.CreateBargainProduct(c, &product.CreateBargainProductRequest{
		ProductId:       data.ProductId,
		Title:           data.Title,
		Stock:           data.Stock,
		StartTime:       data.StartTime,
		StopTime:        data.StopTime,
		MinPrice:        float32(data.BargainMinPrice),
		BargainMaxPrice: data.BargainMaxPrice,
		BargainMinPrice: data.BargainMinPrice,
	})
	fmt.Println(bargainProduct)

	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, bargainProduct)

}

// TODO:商家添加商品
func AddProduct(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.AddProductRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	addProduct, err := handler.AddProduct(c, &product.AddProductRequest{
		MerId:     int64(userid),
		Image:     data.Image,
		StoreName: data.StoreName,
		StoreInfo: data.StoreInfo,
		Keyword:   data.Keyword,
		CateId:    data.CateId,
		Price:     float32(data.Price),
		VipPrice:  float32(data.VipPrice),
		OtPrice:   float32(data.OtPrice),
		Cost:      float32(data.Cost),
		Stock:     data.Stock,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, addProduct)
}

// TODO:创建砍价用户
func CreateBargainUser(c *gin.Context) {
	var data request.CreateBargainUserRequest
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	userRequest, err := handler.CreateBargainUserRequest(c, &product.CreateBargainUserRequest{
		Uid:       int64(id),
		BargainId: data.BargainId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	response.ResponseSuccess(c, userRequest)

}

// TODO:商家添加拼团商品
func AddGroupProduct(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.AddGroupProductRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	groupProduct, err := handler.AddGroupProduct(c, &product.AddGroupProductRequest{
		MerId:     int64(userid),
		ProductId: data.ProductId,
		Title:     data.Title,
		Attr:      data.Attr,
		People:    data.People,
		Info:      data.Info,
		Price:     float32(data.Price),
		Sort:      data.Sort,
		Sales:     data.Sales,
		Stock:     data.Stock,
		StartTime: data.StartTime,
		StopTime:  data.StopTime,
		Cost:      data.Cost,
	})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, groupProduct)
}

// TODO:商家删除拼团商品
func RemoveGroupProduct(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.RemoveGroupProductRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	groupProduct, err := handler.RemoveGroupProduct(c, &product.RemoveGroupProductRequest{
		MerId:     int64(userid),
		ProductId: data.ProductId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, groupProduct)
}

func CartTotalPrice(c *gin.Context) {
	userid := c.GetUint("userId")
	groupProduct, err := handler.CartTotalPrice(c, &product.CartTotalPriceRequest{UserId: int64(userid)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	response.ResponseSuccess(c, groupProduct)
}

func CreateBargainUserHelp(c *gin.Context) {
	var data request.CreateBargainUserHelpRequest
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	help, err := handler.CreateBargainUserHelp(c, &product.CreateBargainUserHelpRequest{
		Uid:           uint32(id),
		BargainId:     uint32(data.BargainId),
		BargainUserId: uint32(data.BargainUserId),
		InviteCode:    data.InviteCode,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, help)
}

// TODO:删除
func MerchantDeletePro(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.MerDeleteProRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	pro, err := handler.MerchantDeletePro(c, &product.MerchantDeleteProRequest{
		MerId:     int64(userid),
		ProductId: data.ProductId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, pro)
}

// TODO:用户发起拼团
func UserGroupBuying(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.UserGroupBuyingRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	buying, err := handler.CreateUserGroup(c, &product.CreateUserGroupRequest{
		Uid: int64(userid),
		Cid: data.Cid,
		Num: data.Num,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, buying)
}

// TODO:商品列表展示
func ProductList(c *gin.Context) {
	list, err := handler.ProductList(c, &product.ProductListRequest{})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, list)
}

func AddSpikeProduct(c *gin.Context) {
	var data request.AddSpikeProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	spikeProduct, err := handler.AddSpikeProduct(c, &product.AddSpikeProductRequest{
		ProductId:  data.ProductId,
		SpikePrice: float32(data.SpikePrice),
		SpikeStock: data.SpikeStock,
		StartTime:  data.StartTime,
		EndTime:    data.EndTime,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, spikeProduct)

}

// TODO:添加优惠卷
func AddCoupon(c *gin.Context) {
	var data request.AddCouponRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	coupon, err := handler.AddCoupon(c, &product.AddCouponRequest{
		Title:       data.Title,
		Integral:    data.Integral,
		CouponPrice: float32(data.CouponPrice),
		UseMinPrice: float32(data.UseMinPrice),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, coupon)
}

// TODO:商家添加优惠卷
func MerAddCouponStore(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.MerAddCouponStoreRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	coupon, err := handler.MerAddCoupon(c, &product.MerAddCouponStoreRequest{
		MerId:     int64(userid),
		Cid:       data.Cid,
		StartTime: data.StartTime,
		StopTime:  data.EndTime,
		Num:       data.Num,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, coupon)
}

// TODO: 商家删除优惠卷
func MerDeleteCouponStore(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.MerDelCouponStoreRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	coupon, err := handler.MerDelCouponStore(c, &product.MerDeleteCouponStoreRequest{
		MerId: int64(userid),
		Cid:   data.Cid,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, coupon)

}

func FlashSale(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.FlashSale
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	sale, err := handler.FlashSale(c, &product.FlashSaleRequest{
		SpikeProductsId: data.SpikeProductsId,
		UserId:          int64(userid),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, sale)
}

// TODO:用户加入拼团
func UserJoinGroup(c *gin.Context) {
	//userid := c.GetUint("userId")
	var data request.UserJoinGroupRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	joinGroup, err := handler.UserJoinGroup(c, &product.UserJoinGroupRequest{
		Uid:            data.UserId,
		Cid:            data.Cid,
		Num:            data.Num,
		InvitationCode: data.InvitationCode,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, joinGroup)
}

// todo 商品信息写入es
func ProductCreateToEs(c *gin.Context) {
	var data request.ToEs
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	es, err := handler.ProductCreateToEs(c, &product.ProductCreateToESRequest{Table: data.TableName})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, es)
}

// todo:商品添加至购物车
func AddCart(c *gin.Context) {
	userId := c.GetUint("userId")
	fmt.Println(userId)
	var data request.AddCartRequest
	cart, err := handler.AddCart(c, &product.AddCartRequest{
		UserId:    int64(userId),
		ProductId: int64(data.ProductId),
		Num:       int64(data.Num),
	})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, cart.Success)
}

// todo 查询es中的商品信息
func ProductSearchToEs(c *gin.Context) {
	var data request.SearchToEs
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	es, err := handler.ProductSearchToEs(c, &product.ProductSearchESRequest{Name: data.Name})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, es)
}

// todo: 移除购物车中商品
func DeleteProductFromCart(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.RemoveFromCartRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	cart, err := handler.DeleteProductFromCart(c, &product.RemoveFromCartRequest{
		UserId:    int64(userId),
		ProductId: int64(data.ProductId),
	})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, cart.Success)
}

// todo: 购物车商品数量修改
func UpdateProductCart(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.UpdateCartRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	cart, err := handler.UpdateProductCart(c, &product.UpdateCartRequest{
		UserId:    int64(userId),
		ProductId: int64(data.ProductId),
		Num:       int64(data.Num),
	})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, cart)
}

// todo: 清空购物车
func ClearCart(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.ClearCart
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	cart, err := handler.ClearCart(c, &product.ClearCartRequest{UserId: int64(userId)})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, cart)
}

// todo: 购物车商品列表
func CartProductList(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.CartProductList
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	list, err := handler.CartProductList(c, &product.CartProductListRequest{UserId: int64(userId)})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, list)
}

// todo: 购物车商品总数量
func CartProductCount(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.CartProductCount
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	count, err := handler.CartProductCount(c, &product.CartProductCountRequest{UserId: int64(userId)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, count)
}

// todo:购物车商品总价
func CartProductTotalPrice(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.CartProductTotalPrice
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	price, err := handler.CartProductTotalPrice(c, &product.CartProductTotalPriceRequest{UserId: int64(userId)})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, price)
}
