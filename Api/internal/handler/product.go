package handler

import (
	"Api/client"
	"context"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品详情
func ProductInfo(ctx context.Context, i *product.ProductInfoRequest) (*product.ProductInfoResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductInfo(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductInfoResponse), nil
}

// CreateBargainProduct 添加砍价商品
func CreateBargainProduct(ctx context.Context, i *product.CreateBargainProductRequest) (*product.CreateBargainProductResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		bargainProduct, err := in.CreateBargainProduct(ctx, i)
		if err != nil {
			return nil, err
		}
		return bargainProduct, nil

	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.CreateBargainProductResponse), nil
}

// TODO：商品推荐
func ProductRecommend(ctx context.Context, i *product.ProductRecommendRequest) (*product.ProductRecommendResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductRecommend(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductRecommendResponse), nil
}

// AddProduct 商家发布商品
func AddProduct(ctx context.Context, i *product.AddProductRequest) (*product.AddProductResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		bargainProduct, err := in.AddProduct(ctx, i)
		if err != nil {
			return nil, err
		}
		return bargainProduct, nil

	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.AddProductResponse), nil

}

// TODO：商品关键字查询
func ProductKeyword(ctx context.Context, i *product.ProductKeywordRequest) (*product.ProductKeywordResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductKeyword(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductKeywordResponse), nil
}

// CreateBargainUserRequest 添加砍价用户
func CreateBargainUserRequest(ctx context.Context, i *product.CreateBargainUserRequest) (*product.CreateBargainUserResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		user, err := in.CreateBargainUser(ctx, i)
		if err != nil {
			return nil, err
		}
		return user, err

	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CreateBargainUserResponse), nil
}

// TODO：商品筛选
func ProductFilter(ctx context.Context, i *product.ProductFilterRequest) (*product.ProductFilterResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductFilter(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductFilterResponse), nil
}

// AddGroupProduct 商家添加拼团商品
func AddGroupProduct(ctx context.Context, i *product.AddGroupProductRequest) (*product.AddGroupProductResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		groupProduct, err := in.AddGroupProduct(ctx, i)
		if err != nil {
			return nil, err
		}
		return groupProduct, err
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.AddGroupProductResponse), nil
}

// CreateBargainUserHelp 帮砍
func CreateBargainUserHelp(ctx context.Context, i *product.CreateBargainUserHelpRequest) (*product.CreateBargainUserHelpResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		help, err := in.CreateBargainUserHelp(ctx, i)
		if err != nil {
			return nil, err
		}
		return help, err

	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CreateBargainUserHelpResponse), nil

}

// TODO：商品分类
func ProductCategory(ctx context.Context, i *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductCategory(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductCategoryResponse), nil
}

// RemoveGroupProduct 删除拼团商品
func RemoveGroupProduct(ctx context.Context, i *product.RemoveGroupProductRequest) (*product.RemoveGroupProductResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		groupProduct, err := in.RemoveGroupProduct(ctx, i)
		if err != nil {
			return nil, err
		}
		return groupProduct, nil
	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.RemoveGroupProductResponse), nil

}

// 商家删除商品
func MerchantDeletePro(ctx context.Context, i *product.MerchantDeleteProRequest) (*product.MerchantDeleteProResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		groupProduct, err := in.MerchantDeletePro(ctx, i)
		if err != nil {
			return nil, err
		}
		return groupProduct, nil

	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.MerchantDeleteProResponse), nil
}

// TODO：添加购物车
func AddToCart(ctx context.Context, i *product.AddToCartRequest) (*product.AddToCartResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.AddToCart(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.AddToCartResponse), nil
}

// TODO：购物车计算总价
func CartTotalPrice(ctx context.Context, i *product.CartTotalPriceRequest) (*product.CartTotalPriceResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.CartTotalPrice(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CartTotalPriceResponse), nil

}

// TODO:用户拼团
func CreateUserGroup(ctx context.Context, i *product.CreateUserGroupRequest) (*product.CreateUserGroupResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		groupProduct, err := in.CreateUserGroup(ctx, i)
		if err != nil {
			return nil, err
		}
		return groupProduct, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CreateUserGroupResponse), nil
}

// TODO: 商品列表展示
func ProductList(ctx context.Context, i *product.ProductListRequest) (*product.ProductListResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.ProductList(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductListResponse), nil
}

// todo:添加秒杀商品
func AddSpikeProduct(ctx context.Context, i *product.AddSpikeProductRequest) (*product.AddSpikeProductResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		spikeProduct, err := in.AddSpikeProduct(ctx, i)
		if err != nil {
			return nil, err
		}
		return spikeProduct, err
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.AddSpikeProductResponse), nil
}

// TODO: 添加优惠卷
func AddCoupon(ctx context.Context, i *product.AddCouponRequest) (*product.AddCouponResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.AddCoupon(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.AddCouponResponse), nil
}

// TODO: 商家添加优惠卷
func MerAddCoupon(ctx context.Context, i *product.MerAddCouponStoreRequest) (*product.MerAddCouponStoreResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.MerAddCouponStore(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.MerAddCouponStoreResponse), nil
}

// TODO: 商家删除优惠卷
func MerDelCouponStore(ctx context.Context, i *product.MerDeleteCouponStoreRequest) (*product.MerDeleteCouponStoreResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.MerDeleteCouponStore(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.MerDeleteCouponStoreResponse), nil
}

func FlashSale(ctx context.Context, i *product.FlashSaleRequest) (*product.FlashSaleResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		sale, err := in.FlashSale(ctx, i)
		if err != nil {
			return nil, err
		}
		return sale, err
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.FlashSaleResponse), nil
}

// todo:用户参与拼团
func UserJoinGroup(ctx context.Context, i *product.UserJoinGroupRequest) (*product.UserJoinGroupResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		sale, err := in.UserJoinGroup(ctx, i)
		if err != nil {
			return nil, err
		}
		return sale, err
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.UserJoinGroupResponse), nil
}

// todo 将商品信息写入es
func ProductCreateToEs(ctx context.Context, i *product.ProductCreateToESRequest) (*product.ProductCreateToESResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		es, err := in.ProductCreateToES(ctx, i)
		if err != nil {
			return nil, err
		}
		return es, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductCreateToESResponse), nil
}

// todo:商品添加至购物车
func AddCart(ctx context.Context, i *product.AddCartRequest) (*product.AddCartResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		cart, err := in.AddCart(ctx, i)
		if err != nil {
			return nil, err
		}
		return cart, nil

	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.AddCartResponse), nil
}

// todo 查询es中的商品信息
func ProductSearchToEs(ctx context.Context, i *product.ProductSearchESRequest) (*product.ProductSearchESResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		es, err := in.ProductSearchES(ctx, i)
		if err != nil {
			return nil, err
		}
		return es, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ProductSearchESResponse), nil
}

// TODO: 移除购物车中商品
func DeleteProductFromCart(ctx context.Context, i *product.RemoveFromCartRequest) (*product.RemoveFromCartResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		cart, err := in.RemoveFromCart(ctx, i)
		if err != nil {
			return nil, err
		}
		return cart, nil

	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.RemoveFromCartResponse), nil
}

// TODO: 修改购物车中商品的数量
func UpdateProductCart(ctx context.Context, i *product.UpdateCartRequest) (*product.UpdateCartResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		cart, err := in.UpdateCart(ctx, i)

		if err != nil {
			return nil, err
		}
		return cart, nil
	})
	if err != nil {
		return nil, err
	}

	return productClient.(*product.UpdateCartResponse), nil
}

// TODO: 清空购物车
func ClearCart(ctx context.Context, i *product.ClearCartRequest) (*product.ClearCartResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		cart, err := in.ClearCart(ctx, i)
		if err != nil {
			return nil, err
		}
		return cart, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.ClearCartResponse), nil
}

// TODO: 购物车商品列表
func CartProductList(ctx context.Context, i *product.CartProductListRequest) (*product.CartProductListResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		list, err := in.CartProductList(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CartProductListResponse), nil
}

// TODO：购物车中商品的总数量
func CartProductCount(ctx context.Context, i *product.CartProductCountRequest) (*product.CartProductCountResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		count, err := in.CartProductCount(ctx, i)
		if err != nil {
			return nil, err
		}
		return count, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CartProductCountResponse), nil
}

// TODO: 购物车中商品总价
func CartProductTotalPrice(ctx context.Context, i *product.CartProductTotalPriceRequest) (*product.CartProductTotalPriceResponse, error) {
	productClient, err := client.ProductClient(ctx, func(ctx context.Context, in product.ProductClient) (interface{}, error) {
		price, err := in.CartProductTotalPrice(ctx, i)
		if err != nil {
			return nil, err
		}
		return price, nil
	})
	if err != nil {
		return nil, err
	}
	return productClient.(*product.CartProductTotalPriceResponse), nil
}
