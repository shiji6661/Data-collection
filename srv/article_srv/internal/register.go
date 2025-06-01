package internal

import (
	"article_srv/proto_article/article"
	server2 "article_srv/server"
	"google.golang.org/grpc"
)

func RegisterArticle(server grpc.ServiceRegistrar) {
	article.RegisterArticleServer(server, server2.ServerArticle{})
}
