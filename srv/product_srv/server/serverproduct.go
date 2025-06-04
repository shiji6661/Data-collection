package server

import (
	"context"
	"product_srv/internal/logic"
	"product_srv/proto_product/product"
)

type ServerProduct struct {
	product.UnimplementedProductServer
}

// todo:商品详情
func (s ServerProduct) ProductInfo(ctx context.Context, in *product.ProductInfoRequest) (*product.ProductInfoResponse, error) {
	res, err := logic.ProductInfo(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:商品推荐
func (s ServerProduct) ProductRecommend(ctx context.Context, in *product.ProductRecommendRequest) (*product.ProductRecommendResponse, error) {
	res, err := logic.ProductRecommend(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:商品关键字查询
func (s ServerProduct) ProductKeyword(ctx context.Context, in *product.ProductKeywordRequest) (*product.ProductKeywordResponse, error) {
	res, err := logic.ProductKeyword(in)

	if err != nil {
		return nil, err
	}
	return res, nil

}

// todo:商品筛选
func (s ServerProduct) ProductFilter(ctx context.Context, in *product.ProductFilterRequest) (*product.ProductFilterResponse, error) {
	res, err := logic.ProductFilter(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:商品分类查询
func (s ServerProduct) ProductCategory(ctx context.Context, in *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	res, err := logic.ProductCategory(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:添加购物车
func (s ServerProduct) AddToCart(ctx context.Context, in *product.AddToCartRequest) (*product.AddToCartResponse, error) {
	res, err := logic.AddToCart(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo: CreateBargainProduct 创建砍价商品
func (s ServerProduct) CreateBargainProduct(ctx context.Context, in *product.CreateBargainProductRequest) (*product.CreateBargainProductResponse, error) {
	bargainProduct, err := logic.CreateBargainProduct(in)
	if err != nil {
		return nil, err
	}
	return bargainProduct, err
}

// todo: AddProduct 商家添加商品
func (s ServerProduct) AddProduct(ctx context.Context, in *product.AddProductRequest) (*product.AddProductResponse, error) {
	bargainProduct, err := logic.AddProduct(in)
	if err != nil {
		return nil, err
	}
	return bargainProduct, nil
}

// todo: CreateBargainUser 创建砍价用户
func (s ServerProduct) CreateBargainUser(ctx context.Context, in *product.CreateBargainUserRequest) (*product.CreateBargainUserResponse, error) {
	user, err := logic.CreateBargainUser(in)
	if err != nil {
		return nil, err
	}
	return user, err
}

// todo: AddGroupProduct 添加拼团商品
func (s ServerProduct) AddGroupProduct(ctx context.Context, in *product.AddGroupProductRequest) (*product.AddGroupProductResponse, error) {
	buyingProduct, err := logic.CreateGroupBuyingProduct(in)
	if err != nil {
		return nil, err
	}
	return buyingProduct, nil
}

// todo: RemoveGroupProduct 删除拼团商品
func (s ServerProduct) RemoveGroupProduct(ctx context.Context, in *product.RemoveGroupProductRequest) (*product.RemoveGroupProductResponse, error) {
	buyingProduct, err := logic.MerRemoveGProduct(in)
	if err != nil {
		return nil, err
	}
	return buyingProduct, nil
}

// todo 购物车总价计算
func (s ServerProduct) CartTotalPrice(ctx context.Context, in *product.CartTotalPriceRequest) (*product.CartTotalPriceResponse, error) {
	buyingProduct, err := logic.CartTotalPrice(in)
	if err != nil {
		return nil, err
	}
	return buyingProduct, nil
}

// todo: CreateBargainUserHelp 创建砍价用户帮助
func (s ServerProduct) CreateBargainUserHelp(ctx context.Context, in *product.CreateBargainUserHelpRequest) (*product.CreateBargainUserHelpResponse, error) {
	help, err := logic.CreateBargainUserHelp(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 删除商品
func (s ServerProduct) MerchantDeletePro(ctx context.Context, in *product.MerchantDeleteProRequest) (*product.MerchantDeleteProResponse, error) {
	help, err := logic.MerDeletePro(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 用户拼团
func (s ServerProduct) CreateUserGroup(ctx context.Context, in *product.CreateUserGroupRequest) (*product.CreateUserGroupResponse, error) {
	help, err := logic.UserCreateBuyingProduct(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 商品列表展示
func (s ServerProduct) ProductList(ctx context.Context, in *product.ProductListRequest) (*product.ProductListResponse, error) {
	help, err := logic.ProductList(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 添加优惠卷
func (s ServerProduct) AddCoupon(ctx context.Context, in *product.AddCouponRequest) (*product.AddCouponResponse, error) {
	help, err := logic.AddCoupon(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 商家添加优惠卷
func (s ServerProduct) MerAddCouponStore(ctx context.Context, in *product.MerAddCouponStoreRequest) (*product.MerAddCouponStoreResponse, error) {
	help, err := logic.MerAddCouponStore(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 商家删除优惠卷
func (s ServerProduct) MerDeleteCouponStore(ctx context.Context, in *product.MerDeleteCouponStoreRequest) (*product.MerDeleteCouponStoreResponse, error) {
	help, err := logic.MerDelCouponStore(in)
	if err != nil {
		return nil, err
	}
	return help, err
}

// todo: 商家添加秒杀商品
func (s ServerProduct) AddSpikeProduct(ctx context.Context, in *product.AddSpikeProductRequest) (*product.AddSpikeProductResponse, error) {
	spikeProduct, err := logic.AddSpikeProduct(in)
	if err != nil {
		return nil, err
	}
	return spikeProduct, err
}

func (s ServerProduct) FlashSale(ctx context.Context, in *product.FlashSaleRequest) (*product.FlashSaleResponse, error) {
	sale, err := logic.FlashSale(in)
	if err != nil {
		return nil, err
	}
	return sale, err
}

// todo:用户参与拼团
func (s ServerProduct) UserJoinGroup(ctx context.Context, in *product.UserJoinGroupRequest) (*product.UserJoinGroupResponse, error) {
	sale, err := logic.UserJoinGroup(in)
	if err != nil {
		return nil, err
	}
	return sale, err
}

// todo 将商品信息写入ES
func (s ServerProduct) ProductCreateToES(ctx context.Context, in *product.ProductCreateToESRequest) (*product.ProductCreateToESResponse, error) {
	es, err := logic.ProductCreateToES(in)
	if err != nil {
		return nil, err
	}
	return es, err
}

// todo 查询ES中的商品信息
func (s ServerProduct) ProductSearchES(ctx context.Context, in *product.ProductSearchESRequest) (*product.ProductSearchESResponse, error) {
	es, err := logic.ProductSearchToEs(in)
	if err != nil {
		return nil, err
	}
	return es, nil
}
