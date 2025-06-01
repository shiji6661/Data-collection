package request

// TODO: 商品详情
type ProductInfo struct {
	Name string `form:"name" json:"name"  binding:"required"`
}

// TODO: 商品关键字查询
type Keyword struct {
	Keyword string `form:"keyword" json:"keyword"`
}

// TODO: 商品筛选
type Filter struct {
	Min_price  float64 `form:"min_Price" json:"min_Price"`
	Max_price  float64 `form:"max_Price" json:"max_Price"`
	Cate_id    string  `form:"cate_Id" json:"cate_Id"`
	Sales      int64   `form:"sales" json:"sales"`
	Is_show    int64   `form:"is_Show" json:"is_Show"`
	Is_hot     int64   `form:"is_Hot" json:"is_Hot"`
	Is_benefit int64   `form:"is_Benefit" json:"is_Benefit"`
	Is_best    int64   `form:"is_Best" json:"is_Best"`
	Is_new     int64   `form:"is_New" json:"is_New"`
	Is_seckill int64   `form:"is_Seckill" json:"is_Seckill"`
	Is_bargain int64   `form:"is_Bargain" json:"is_Bargain"`
	Name       string  `form:"name" json:"name"`
}

// TODO: 商品分类
type ProductCategory struct {
	CateId string `form:"CateId" json:"CateId" binding:"required"`
}

// TODO: 添加购物车
type AddToCart struct {
	ProductId int64 `form:"productId" json:"productId" binding:"required"`
	Num       int64 `form:"num" json:"num" binding:"required"`
}

type CreateBargainProductRequest struct {
	ProductId       int64   `json:"product_id" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	Stock           int64   `json:"stock" binding:"required"`
	StartTime       string  `json:"start_time" binding:"required"`
	StopTime        string  `json:"stop_time" binding:"required"`
	MinPrice        float64 `json:"min_price" binding:"required"`
	BargainMaxPrice int64   `json:"bargain_max_price" binding:"required"`
	BargainMinPrice int64   `json:"bargain_min_price" binding:"required"`
}

// TODO:商家发布商品
type AddProductRequest struct {
	Image     string  `form:"image" binding:"required"`
	StoreName string  `form:"store_name" binding:"required"`
	StoreInfo string  `form:"store_info" binding:"required"`
	Keyword   string  `form:"keyword" binding:"required"`
	CateId    string  `form:"cate_id" binding:"required"`
	Price     float64 `form:"price" binding:"required"`
	VipPrice  float64 `form:"vip_price" binding:"required"`
	OtPrice   float64 `form:"ot_price" binding:"required"`
	Cost      float64 `form:"cost" binding:"required"`
	Stock     int64   `form:"stock" binding:"required"`
}

type CreateBargainUserRequest struct {
	BargainId int64 `form:"bargain_id" json:"bargain_id" binding:"required"`
}

// TODO:商家添加拼团商品
type AddGroupProductRequest struct {
	ProductId int64   `form:"product_id" binding:"required"`
	Title     string  `form:"title" binding:"required"`
	Attr      string  `form:"attr" binding:"required"`
	People    int64   `form:"people" binding:"required"`
	Info      string  `form:"info" binding:"required"`
	Price     float64 `form:"price" binding:"required"`
	Sort      int64   `form:"sort" binding:"required"`
	Sales     int64   `form:"sales" binding:"required"`
	Stock     int64   `form:"stock" binding:"required"`
	StartTime string  `form:"start_time" binding:"required"`
	StopTime  string  `form:"stop_time" binding:"required"`
	Cost      int64   `form:"cost" binding:"required"`
}

// TODO:商家删除拼团商品
type RemoveGroupProductRequest struct {
	ProductId int64 `form:"product_id" binding:"required"`
}

// TODO:帮砍价
type CreateBargainUserHelpRequest struct {
	BargainId     int64 `json:"bargain_id" binding:"required"`
	BargainUserId int64 `json:"bargain_user_id" binding:"required"`
}

// TODO:删除商品
type MerDeleteProRequest struct {
	ProductId int64 `form:"product_id" binding:"required"`
}

// TODO:用户发起拼团
type UserGroupBuyingRequest struct {
	Cid int64 `form:"cid" binding:"required"`
	Num int64 `form:"num" binding:"required"`
}

// TODO:添加优惠卷
type AddCouponRequest struct {
	Title       string  `form:"title" binding:"required"`
	Integral    int64   `form:"integral"`
	CouponPrice float64 `form:"coupon_price" binding:"required"`
	UseMinPrice float64 `form:"use_min_price"`
}

// TODO:商家添加优惠卷
type MerAddCouponStoreRequest struct {
	Cid       int64  `form:"cid" binding:"required"`
	StartTime string `form:"start_time" binding:"required"`
	EndTime   string `form:"end_time" binding:"required"`
	Num       int64  `form:"num" binding:"required"`
}

// TODO:商家删除优惠卷
type MerDelCouponStoreRequest struct {
	Cid int64 `form:"cid" binding:"required"`
}

// AddSpikeProduct todo:添加秒杀商品
type AddSpikeProduct struct {
	ProductId  int64   `json:"product_id" binding:"required"`
	SpikePrice float64 `json:"spike_price" binding:"required"`
	SpikeStock int64   `json:"spike_stock" binding:"required"`
	StartTime  string  `json:"start_time" binding:"required"`
	EndTime    string  `json:"end_time" binding:"required"`
}

type FlashSale struct {
	SpikeProductsId int64 `json:"spike_products_id" binding:"required"`
}

// TODO:用户参与拼团
type UserJoinGroupRequest struct {
	Cid            int64  `form:"cid" binding:"required"`
	Num            int64  `form:"num" binding:"required"`
	InvitationCode string `form:"invitation_code"`
}
