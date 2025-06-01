package global

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// todo:全局变量放置
var (
	DB     *gorm.DB
	Rdb    *redis.Client
	Es     *elasticsearch.Client
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
	CTX    = context.Background()
)

const SEND = "send_kill"

const Database = "article"

const CollectionName = "article_category"

const Article = "article"

const ArticleContent = "article_content"

// 假设每个商品在购物车中的最大数量为 99
const Shopping_cart_quantity_limit = 99

const ArticleComment = "article_comment"

// todo 会员积分规则
const (
	POINTS_PER_CONSUMPTION = 1    // 每消费20元获得1积分
	POINTS_PER_INVITATION  = 20   // 每邀请一人注册获得20积分
	CONSUMPTION_UNIT       = 20.0 // 消费单位（元）
	MAX_MEMBER_LEVEL       = 5    // 最高会员等级
	POINTS_PER_LEVEL       = 2000 // 每级所需积分
)
