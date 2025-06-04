package router

import (
	"Api/internal/middleware"
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadProduct(r *gin.Engine) {
	r.Use(middleware.Logger())
	product := r.Group("/product")
	{
		product.POST("/create", trigger.CreateBargainProduct) // todo:用户手机号注册

		product.POST("/list", trigger.ProductList) // todo:商品列表展示

		product.POST("/add/coupon", trigger.AddCoupon) // todo:添加优惠卷

		product.GET("/create/to/es", trigger.ProductCreateToEs) //todo 将商品信息写入ES

		product.GET("/search/es", trigger.ProductSearchToEs) //todo 查询ES中的商品信息

		// 用户中间件
		product.Use(pkg.JWTAuth("2209A"))
		product.POST("/createBargainUser", trigger.CreateBargainUser) // todo:用户参与砍价

		product.POST("/info", trigger.ProductInfo) //todo:商品详情

		product.POST("/recommend", trigger.ProductRecommend) //todo:商品推荐

		product.POST("/keyword", trigger.ProductKeyword) //todo:商品关键字查询

		product.POST("/filter", trigger.ProductFilter) //todo:商品筛选

		product.POST("/category", trigger.ProductCategory) //todo:商品分类查询

		product.POST("/add/cart", trigger.AddToCart) //todo:添加购物车

		product.POST("/cart/total/price", trigger.CartTotalPrice) //todo:购物车计算总价

		product.POST("/createBargainUserHelp", trigger.CreateBargainUserHelp) //todo:帮砍

		product.POST("/user/group/buying", trigger.UserGroupBuying) //todo:用户拼团

		product.POST("/flash/sale", trigger.FlashSale)

		product.POST("/user/join/group", trigger.UserJoinGroup) //todo:用户加入拼团
	}
	productMerchant := r.Group("/product")
	{
		// 商户中间件
		productMerchant.Use(pkg.JWTAuth("Merchant"))
		productMerchant.POST("/add/product", trigger.AddProduct) // todo:商户添加商品

		productMerchant.POST("/createBargain", trigger.CreateBargainProduct) // todo:商户创建砍价活动

		productMerchant.POST("/add/group/product", trigger.AddGroupProduct) // todo:商家添加拼团商品

		productMerchant.POST("/remove/group/product", trigger.RemoveGroupProduct) // todo:商家删除拼团商品
		productMerchant.POST("/delete/product", trigger.MerchantDeletePro)        // todo:商家删除商品
		productMerchant.POST("/add/spike/product", trigger.AddSpikeProduct)       //todo:商家添加秒杀商品

		productMerchant.POST("/add/coupon/store", trigger.MerAddCouponStore) // TODO: 商家添加优惠卷

		productMerchant.POST("/delete/coupon/store", trigger.MerDeleteCouponStore) // TODO: 商家删除优惠卷
	}

}
